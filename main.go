package main

import (
	"build-dotenv/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Debug("Starting...")
	cmd.Execute()
}
