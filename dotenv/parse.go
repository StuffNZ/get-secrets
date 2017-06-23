package dotenv

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

// AddFromString TODO
func (s *DotEnvs) AddFromString(path string, body string) error {
	if path == "" {
		log.Error("Empty path!")
		return fmt.Errorf("Empty path provided")
	}
	if body == "" {
		log.Error("Empty body!")
		return fmt.Errorf("Empty body from %#v", path)
	}

	log.WithFields(log.Fields{"path": path, "body": body}).Debug("Parsing dotenv file...")
	s.env[path] = BodyEnv{path: path, body: body, env: s.parseEnv(body)}

	return nil
}

// Join TODO
// func (s *DotEnvs) Join() map[string]string {
// }

func (s *DotEnvs) parseEnv(body string) gotenv.Env {
	env := gotenv.Parse(strings.NewReader(body))
	return env
}
