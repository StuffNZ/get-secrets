package main

import (
	"fmt"
	"os"

	"bitbucket.org/mexisme/get-secrets/config"
	"bitbucket.org/mexisme/get-secrets/dotenv"
	execish "bitbucket.org/mexisme/get-secrets/exec"
	s3ish "bitbucket.org/mexisme/get-secrets/files/s3"
	urlish "bitbucket.org/mexisme/get-secrets/files/s3/s3url"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	config.SetVersion(Version)
}

func main() {
	log.Debug("Starting...")

	dotenvs := dotenv.New()

	s3url := urlish.New().WithURL(viper.GetString("s3.path"))
	s3 := s3ish.New().WithSource(s3url)
	s3lists, _ := s3.List()

	s3.ReadList(s3lists, dotenvs.AddFromString)

	runner := execish.New().WithEnviron(os.Environ()).WithDotEnvs(dotenvs)
	if len(os.Args) > 1 {
		log.Panic(runner.WithCommand(os.Args[1:]).Exec())
	}

	// log.WithFields(log.Fields{"env": runner.CombineEnvs()}).Info("DONE")
	fmt.Println("# No command provided to execute")
	for _, envLine := range runner.CombineEnvs() {
		fmt.Printf("export %s\n", envLine)
	}
}
