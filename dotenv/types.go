package dotenv

import (
	"fmt"
	// "net/url"
	"strings"

	log "github.com/sirupsen/logrus"
	// "github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

// BodyEnv TODO
type BodyEnv struct {
	path string
	body string
	env  gotenv.Env
}

// DotEnvs TODO
type DotEnvs map[string]BodyEnv

// New TODO
func New() DotEnvs {
	return make(DotEnvs)
}

// AddFromString TODO
func (s DotEnvs) AddFromString(path string, body string) error {
	if path == "" {
		log.Error("Empty path!")
		return fmt.Errorf("Empty path provided")
	}
	if body == "" {
		log.Error("Empty body!")
		return fmt.Errorf("Empty body from %#v", path)
	}

	log.WithFields(log.Fields{"path": path, "body": body}).Debug("Parsing dotenv file...")

	s[path] = BodyEnv{path: path, body: body, env: s.parseEnv(body)}

	return nil
}

// Join TODO
// func (s *DotEnvs) Join() map[string]string {
// }

func (s *DotEnvs) parseEnv(body string) gotenv.Env {
	return gotenv.Parse(strings.NewReader(body))
}
