package main

import (
	"build-dotenv/config"

	log "github.com/sirupsen/logrus"
	// "github.com/spf13/viper"
	// "build-dotenv/cmd"
)

func init() {
	config.ImportMe()
}

func main() {
	log.Debug("Starting...")
}
