package config

import (
	"bitbucket.org/mexisme/get-secrets/config/logging"
	"bitbucket.org/mexisme/get-secrets/config/settings"
)

//nolint:gochecknoglobals
var (
	// We don't want to try to reinitialise the config more than once
	initConfigDone = false
	// We don't want to try to reinitialise the logging more than once
	logConfigDone = false
)

//nolint:gochecknoinits
func init() {
	LoggingConfig()
}

// AddConfigItems passes the configItems through to config.AddConfigItems()
func AddConfigItems(configItems []string) {
	// Need to ensure the system has been configured at least once!
	readConfig() // TODO: Viper dynamically reads -- this may not be needed.
	settings.AddConfigItems(configItems)
}

func readConfig() {
	if !initConfigDone {
		settings.ReadConfig()

		initConfigDone = true
	}
}

func LoggingConfig() {
	readConfig()

	// This should make it safe to rerun a few times
	if !logConfigDone {
		logging.Configure()

		logConfigDone = true
	}
}
