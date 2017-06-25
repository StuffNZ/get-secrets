package settings

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"os"
)

var initConfigDone = false

// ReadConfig uses Viper to read the configuration from .secrets.* files or Env Vars
// TODO:  list config items
func ReadConfig() {
	viper.SetEnvPrefix("secrets")
	viper.BindEnv("debug")
	viper.BindEnv("base")
	viper.BindEnv("app")

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
