package main

import (
	log "github.com/Sirupsen/logrus"
	s3ish "bitbucket.org/mexisme/build-dotenv/files/s3"
	// "build-dotenv/cmd"
)

func main() {
	log.Debug("Starting...")

	s3lists, _ := (&s3ish.Source{Url: "s3://kiwiops-ecs-staging-env/stuff-brightcove-video-service"}).Init().List()
	log.WithFields(log.Fields{"list": s3lists}).Info()
}
