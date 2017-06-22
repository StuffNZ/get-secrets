package main

import (
	"build-dotenv/config"
	s3ish "build-dotenv/files/s3"
	urlish "build-dotenv/files/s3/s3url"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	// "build-dotenv/cmd"
)

func init() {
	config.ImportMe()
}

// EnvRead TODO
type EnvRead struct {
	bodies []string
	envs   []gotenv.Env
}

var env EnvRead

func main() {
	log.Debug("Starting...")
	liveTest()
}

func liveTest() {
	env := New()

	s3url := urlish.New().WithURL(viper.GetString("s3.path"))
	s3 := s3ish.New().WithSource(s3url)
	s3lists, err := s3.List()
	if err != nil {
		log.Panic(err)
	}
	// s3lists = append(s3lists, "")
	// s3lists = append(s3lists, "lol")
	// s3lists, _ := (&s3ish.Source{URLString: "s3://kiwiops-ecs-staging-env/stuff-brightcove-video-service"}).Init().List()
	log.WithFields(log.Fields{"list": s3lists}).Info()

	// s3lists, _ := (&s3ish.Source{URLString: "s3://kiwiops-ecs-staging-env/stuff-brightcove-video-service"}).Init().List()
	if errs := s3.ReadList(s3lists, env.readCallback); errs == nil {
		log.WithFields(log.Fields{"envs": env.envs}).Info("Have parsed all the envs!")
	} else {
		log.Panic(errs)
	}
}

// New TODO
func New() *EnvRead {
	return &EnvRead{
		bodies: make([]string, 0),
		envs:   make([]gotenv.Env, 0),
	}
}

func (s *EnvRead) readCallback(path, body string) error {
	if path == "" {
		log.Error("Empty path!")
		return fmt.Errorf("Empty path")
	}
	if body == "" {
		log.Error("Empty body!")
		return fmt.Errorf("Empty body")
	}

	env := gotenv.Parse(strings.NewReader(body))
	s.bodies = append(s.bodies, body)
	s.envs = append(s.envs, env)

	log.WithFields(log.Fields{"path": path, "env": env}).Debug(body)
	return nil
}
