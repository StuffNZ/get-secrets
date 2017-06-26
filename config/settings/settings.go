/*
Package settings manages reading config from .secrets file or env vars.

Supported Settings

- debug ($SECRETS_DEBUG) -- enables debug mode.

- base ($SECRETS_BASE) -- sets the base dir for reading/writing config files and .env files

- application.name ($APPLICATION_NAME) -- sets the app name for logging purposes

Note: other packages may add other settings.
*/
package settings

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"os"
	"strings"
)

var initConfigDone = false

// ReadConfig uses Viper to read the configuration from .secrets.* files or Env Vars
// TODO:  list config items
func ReadConfig() {
	viper.SetEnvPrefix("secrets")
	viper.BindEnv("debug")
	viper.BindEnv("base")

	viper.BindEnv("application.name", "APPLICATION_NAME")

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
		// panic(fmt.Errorf("Fatal error config file: %s \n", err))
		log.WithFields(log.Fields{"config_file": viper.ConfigFileUsed()}).Warn(err)
	}
}

// AddConfigItems adds a new configuration item, and makes it overridable by env vars
func AddConfigItems(configItems []string) {
	for _, item := range configItems {
		viper.BindEnv(item)
	}
}
