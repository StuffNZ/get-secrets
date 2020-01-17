/*
Package settings manages reading config from .secrets file or env vars.

Supported Settings

- debug ($SECRETS_DEBUG) -- enables debug mode.

- base ($SECRETS_BASE) -- sets the base dir for reading/writing config files and .env files

- dotenv.skip ($SKIP_SECRETS) -- skips the code which reads .env files from S3

- application.name ($APPLICATION_NAME) -- sets the app name for logging purposes

- application.environment ($ENVIRONMENT) -- sets the app name for logging purposes

Note: other packages may add other settings.
*/
package settings

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"os"
	"strings"
)

// ReadConfig uses Viper to read the configuration from .secrets.* files or Env Vars
// TODO:  list config items
func ReadConfig() {
	viper.SetEnvPrefix("secrets")
	// nolint: gosec
	{
		// NOTE: We wrap these statements into a block to support the 'nolint' above
		//       This is because BindEnv() can return an error if no args are provided:
		if err := viper.BindEnv("debug"); err != nil {
			log.WithField("Error", err).Panic(`When viper.BindEnv("debug")`)
		}
		if err := viper.BindEnv("base"); err != nil {
			log.WithField("Error", err).Panic(`When viper.BindEnv("base")`)
		}

		if err := viper.BindEnv("dotenv.skip", "SKIP_SECRETS"); err != nil {
			log.WithField("Error", err).Panic(`When viper.BindEnv("dotenv.skip")`)
		}

		if err := viper.BindEnv("application.name", "APPLICATION_NAME"); err != nil {
			log.WithField("Error", err).Panic(`When viper.BindEnv("application.name")`)
		}
		if err := viper.BindEnv("application.environment", "ENVIRONMENT"); err != nil {
			log.WithField("Error", err).Panic(`When viper.BindEnv("application.environment")`)
		}
	}

	// This means any "." chars in a FQ config name will be replaced with "_"
	// e.g. "sentry.dsn" --> "$SECRETS_SENTRY_DSN" instead of "$SECRETS_SENTRY.DSN"
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName(".secrets")
	viper.AddConfigPath("$HOME")

	if os.Getenv("SECRETS_BASE") != "" {
		viper.AddConfigPath("$SECRETS_BASE/")
	}

	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err == nil {
		log.WithFields(log.Fields{"config_file": viper.ConfigFileUsed()}).Debug("Using file")
	} else {
		log.WithFields(log.Fields{"config_file": viper.ConfigFileUsed()}).Error(err)
	}
}

// AddConfigItems adds a new configuration item, and makes it overridable by env vars
func AddConfigItems(configItems []string) {
	for _, item := range configItems {
		if err := viper.BindEnv(item); err != nil {
			log.WithFields(log.Fields{"Error": err, "ConfigKey": item}).Panic("When viper.BindEnv($ConfigKey)")
		} // nolint: gosec
	}
}
