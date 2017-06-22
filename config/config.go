package config

import (
	"os"
	"time"

	"github.com/evalphobia/logrus_sentry"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	initConfigDone = false
	logConfigDone  = false
)

func init() {
	readConfig()
	configLogging()
}

// ImportMe is to allow other packages to easily depend on this one,
// since most of the important logic is in init()
func ImportMe() {
}

// AddConfigItems TODO
func AddConfigItems(configItems []string) {
	readConfig()
	for _, item := range configItems {
		viper.BindEnv(item)
	}
}

func readConfig() {
	// This should make it safe to rerun a few times
	if initConfigDone {
		return
	}

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

	initConfigDone = true
}

func configLogging() {
	// This should make it safe to rerun a few times
	if logConfigDone {
		return
	}
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug mode enabled")
	}

	if sentryDsn := viper.GetString("sentry.dsn"); sentryDsn != "" {
		log.WithFields(log.Fields{"sentry.dsn": sentryDsn}).Debug("Configuring connection to Sentry.io")
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
			// It seems as if the default 100ms is too short:
			hook.Timeout = 1 * time.Second
			log.AddHook(hook)

			log.Debug("Sentry enabled")

		} else {
			log.Warn(err)
		}
	}

	// viper.Debug()

	logConfigDone = true
}
