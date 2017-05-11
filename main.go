package main

import (
	s3ish "build-dotenv/files/s3"
	log "github.com/Sirupsen/logrus"
	// "build-dotenv/cmd"
)

func main() {
	log.Debug("Starting...")

	s3lists, _ := (&s3ish.Source{Base: "s3://kiwiops-ecs-staging-env/stuff-brightcove-video-service"}).Init().List()
	log.WithFields(log.Fields{"list": s3lists}).Info()
}
