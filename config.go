package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_sentry"
	"github.com/spf13/viper"
	"os"
)

func init() {
	readConfig()
	configLogging()
}

func readConfig() {
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
		log.WithFields(log.Fields{"config_file": viper.ConfigFileUsed()}).Info("Using file")

	} else {
		// panic(fmt.Errorf("Fatal error config file: %s \n", err))
		log.WithFields(log.Fields{"config_file": viper.ConfigFileUsed()}).Warn(err)
	}
}

func configLogging() {
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug mode enabled")
	}

	if sentryDsn := viper.GetString("sentry.dsn"); sentryDsn != "" {
		tags := map[string]string{
			"app": viper.GetString("app"),
		}
		levels := []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
		}
		if hook, err := logrus_sentry.NewWithTagsSentryHook(sentryDsn, tags, levels); err == nil {
			// Set the Sentry "release" version, dep on the setting in the config:
			for _, releaseKey := range []string{"sentry.release", "version"} {
				if sentryRelease := viper.GetString(releaseKey); sentryRelease != "" {
					log.WithFields(log.Fields{releaseKey: sentryRelease}).Debug()
					hook.SetRelease(sentryRelease)
					break
				}
			}
			hook.StacktraceConfiguration.Enable = true
			log.AddHook(hook)

			log.WithFields(log.Fields{"sentry.dsn": sentryDsn}).Debug("Sentry enabled")

		} else {
			log.Warn(err)
		}
	}

	// viper.Debug()
}
