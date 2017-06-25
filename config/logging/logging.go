package logging

import (
	"time"

	"bitbucket.org/mexisme/get-secrets/version"
	"github.com/evalphobia/logrus_sentry"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configure set-ups the Logrus library -- debug mode, etc
// Currently set-up via Viper
func Configure() {
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug mode enabled")
	}

	if sentryDsn := viper.GetString("sentry.dsn"); sentryDsn != "" {
		setupSentry(sentryDsn)
	}

	// viper.Debug()
}

func setupSentry(sentryDsn string) {
	log.WithFields(log.Fields{"sentry.dsn": sentryDsn}).Debug("Configuring connection to Sentry.io")

	// TODO: Meta-tag for environment
	// Some meta-tags
	tags := map[string]string{
		// TODO: Pick a better name, as this maps to "SECRETS_APP":
		"app": viper.GetString("app"),
	}

	// Sentry will only log for messages of the following severity:
	levels := []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
	}

	// Hook Sentry into Logrus:
	if hook, err := logrus_sentry.NewWithTagsSentryHook(sentryDsn, tags, levels); err == nil {
		// Set the Sentry "release" version:
		log.WithFields(log.Fields{"release": version.Release}).Debug("Setting release version in Sentry")
		hook.SetRelease(version.Release)

		hook.StacktraceConfiguration.Enable = true

		// It seems as if the default 100ms is too short:
		hook.Timeout = 1 * time.Second

		// Now, add it into the Logrus hook-chain
		log.AddHook(hook)

		log.Debug("Sentry enabled")

	} else {
		log.Warn(err)
	}
}
