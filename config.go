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

	if sentry_dsn := viper.GetString("sentry.dsn"); sentry_dsn != "" {
		tags := map[string]string{
			"app": viper.GetString("app"),
		}
		levels := []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
		}
		if hook, err := logrus_sentry.NewWithTagsSentryHook(sentry_dsn, tags, levels); err == nil {
			// Set the Sentry "release" version, dep on the setting in the config:
			for _, release_key := range []string{"sentry.release", "version"} {
				if sentry_release := viper.GetString(release_key); sentry_release != "" {
					log.WithFields(log.Fields{release_key: sentry_release}).Debug()
					hook.SetRelease(sentry_release)
					break
				}
			}
			hook.StacktraceConfiguration.Enable = true
			log.AddHook(hook)

			log.WithFields(log.Fields{"sentry.dsn": sentry_dsn}).Debug("Sentry enabled")

		} else {
			log.Warn(err)
		}
	}

	// viper.Debug()
}
