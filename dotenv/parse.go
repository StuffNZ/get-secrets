package dotenv

import (
	"fmt"
	"sort"
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

// Combine TODO
func (s *DotEnvs) Combine() map[string]string {
	joinedEnv := make(gotenv.Env)

	for _, path := range s.sortedPaths() {
		s.mergeEnv(s.env[path].env, joinedEnv)
	}

	return joinedEnv
}

func (s *DotEnvs) mergeEnv(from, to map[string]string) {
	for k, v := range from {
		log.WithFields(log.Fields{"key": k, "val": v}).Debug()
		to[k] = v
	}
}

func (s *DotEnvs) sortedPaths() []string {
	names := make([]string, 0)

	for name := range s.env {
		names = append(names, name)
	}
	sort.Strings(names)

	return names
}

func (s *DotEnvs) parseEnv(body string) gotenv.Env {
	env := gotenv.Parse(strings.NewReader(body))
	return env
}
