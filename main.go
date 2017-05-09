package main

import (
	log "github.com/Sirupsen/logrus"
	s3ish "build-dotenv/files/s3"
	// "build-dotenv/cmd"
)

func main() {
	log.Debug("Starting...")

	if s3list, err := (&s3ish.Source{Url: "s3://kiwiops-ecs-staging-env/stuff-brightcove-video-service"}).Init(); err != nil {
		log.Fatal(err)

	} else {
		s3lists, _ := s3list.List()
		log.WithFields(log.Fields{"list": s3lists}).Info()
	}
}
