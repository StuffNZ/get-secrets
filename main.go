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
	"bitbucket.org/mexisme/get-secrets/env"
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
	// When any other part of the app panics, we'd prefer to give them a "friendlier" face
	defer func() {
		if recoveryErr := recover(); recoveryErr != nil {
			// TODO: Re-deliver the stack-trace when debugging
			// TODO: Change to using github.com/go-errors/errors to capture stack-traces (instead of panic()!)
			log.WithFields(log.Fields{"Err": recoveryErr}).Debug("Panic captured")
			os.Exit(1)
		}
	}()

	appName, appEnv := viper.GetString("application.name"), viper.GetString("application.environment")
	log.Infof("Preparing to run app %#v in env %#v", appName, appEnv)

	dotenvs := dotenv.New()

	if viper.GetBool("dotenv.skip") {
		log.Info("Not getting .env secrets due to configuration")

	} else {
		s3Path := viper.GetString("s3.dotenv_path")
		s3url := urlish.New().WithURL(s3Path)
		log.Infof("S3 .env Base path = %#v (%#v)", s3Path, s3url)

		s3 := s3ish.New().WithSource(s3url)

		{
			//
			defer log.Warn("NB: Did you perhaps intend to enable 'dotenv.skip' ($SKIP_SECRETS)?")

			s3lists, err := s3.List()
			if err != nil {
				panicErrs(err)
			}

			if err := s3.ReadListToCallback(s3lists, dotenvs.AddFromString); err != nil {
				panicErrs(err)
			}
		}
	}

	envs := env.New().WithOsEnviron(os.Environ()).WithDotEnvs(dotenvs)

	if len(os.Args) > 1 {
		runner := execish.New().WithEnvs(envs)
		panicErrs(runner.WithCommand(os.Args[1:]).Exec())
	}

	fmt.Println("")
	fmt.Println("# No command provided to execute")
	for _, envLine := range envs.Combine() {
		fmt.Printf("export %s\n", envLine)
	}
}

func panicErrs(err error) {
	if merr, ok := err.(*multierror.Error); ok {
		for _, anErr := range merr.Errors {
			log.Panic(anErr)
		}
		// log.Error("Multiple errors from s3.ReadListToCallback()")

	} else {
		log.Panic(err)
	}
}
