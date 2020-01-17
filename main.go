/*
Package main is the main loop.

Note the one additional config field:

- s3.dotenv_path ($SECRETS_S3_DOTENV_PATH) -- defines the dir where to read all .env files from
*/
package main

import (
	"fmt"
	"os"

	"bitbucket.org/mexisme/get-secrets/dotenv"
	"bitbucket.org/mexisme/get-secrets/errors"
	execish "bitbucket.org/mexisme/get-secrets/exec"
	"github.com/mexisme/multiconfig"
	"github.com/mexisme/multiconfig/env"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// When any other part of the app panics, we'd prefer to give them a "friendlier" face
	defer errors.Recovery()

	appName, appEnv := viper.GetString("application.name"), viper.GetString("application.environment")
	log.Infof("Preparing to run app %#v in env %#v", appName, appEnv)

	secrets := multiconfig.New()

	if viper.GetBool("dotenv.skip") {
		log.Info("Not getting .env secrets due to configuration")
	} else {
		dotenv.ReadFromS3(secrets)
	}

	// Add the SSM / KMS stuff in here?
	secrets.AddItem(env.New().FromOsEnviron())
	envs := dotenv.EnvMerge(secrets)

	const argsCount = 1
	if len(os.Args) > argsCount {
		runner := execish.New().WithEnvs(envs)
		errors.PanicOnErrors(runner.WithCommand(os.Args[argsCount:]).Exec())
	}

	log.Info("# No command provided to execute")

	osEnviron, err := envs.ToOsEnviron()
	if err != nil {
		errors.PanicOnErrors(err)
	}

	for _, envLine := range osEnviron {
		fmt.Printf("export %s\n", envLine)
	}
}
