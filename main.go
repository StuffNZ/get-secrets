/*
Package main is the main loop.

Note the one addiional config field:

- s3.dotenv_path ($SECRETS_S3_DOTENV_PATH) -- defines the dir where to read all .env files from
*/
package main

import (
	"fmt"
	"os"

	"bitbucket.org/mexisme/get-secrets/config"
	"bitbucket.org/mexisme/get-secrets/dotenv"
	execish "bitbucket.org/mexisme/get-secrets/exec"
	s3ish "bitbucket.org/mexisme/get-secrets/files/s3"
	urlish "bitbucket.org/mexisme/get-secrets/files/s3/s3url"

	"github.com/hashicorp/go-multierror"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	config.ImportMe()
}

func main() {
	appName, appEnv := appDetails()
	log.Infof("Starting %v in %v", appName, appEnv)

	dotenvs := dotenv.New()

	if !viper.GetBool("dotenv.skip") {
		s3Path := viper.GetString("s3.dotenv_path")
		s3url := urlish.New().WithURL(s3Path)
		log.Infof("S3 .env Base path = %#v (%#v)", s3Path, s3url)

		s3 := s3ish.New().WithSource(s3url)
		s3lists, err := s3.List()
		if err != nil {
			panicErrs(err)
		}

		if err := s3.ReadListToCallback(s3lists, dotenvs.AddFromString); err != nil {
			panicErrs(err)
		}
	}

	runner := execish.New().WithEnviron(os.Environ()).WithDotEnvs(dotenvs)
	if len(os.Args) > 1 {
		panicErrs(runner.WithCommand(os.Args[1:]).Exec())
	}

	fmt.Println("# No command provided to execute")
	for _, envLine := range runner.CombineEnvs() {
		fmt.Printf("export %s\n", envLine)
	}
}
func appDetails() (string, string) {
	appName, appEnv := viper.GetString("application.name"), viper.GetString("application.environment")
	if appName == "" {
		appName = "(UNKNOWN APP)"
	}
	if appEnv == "" {
		appEnv = "(UNKNOWN ENVIRONMENT)"
	}

	return appName, appEnv
}

func panicErrs(err error) {
	if merr, ok := err.(*multierror.Error); ok {
		for _, anErr := range merr.Errors {
			log.Error(anErr)
		}
		log.Panic("Multiple errors from s3.ReadListToCallback()")
	}

	log.Panic(err)
}
