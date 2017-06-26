package exec

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

// CombineEnvs merges the provided os.Environ []string with the provided .env KV maps into a new os.Environ []string
func (s *Details) CombineEnvs() []string {
	newEnvs := s.envToMap(s.env)
	dotEnvs := s.dotEnvs.Combine()
	log.WithFields(log.Fields{"env": newEnvs, "dotEnv": dotEnvs}).Debug("Combining envs...")

	s.mergeEnv(dotEnvs, newEnvs)

	log.WithFields(log.Fields{"new": newEnvs}).Debug("Combined envs")

	return s.mapToEnv(newEnvs)
}

/*
Merge the first map *into* the second

Because maps are implicit references, updating the second map *updates* the caller's map in-place
*/
func (s *Details) mergeEnv(from, to map[string]string) {
	for k, v := range from {
		to[k] = v
	}
}

func (s *Details) envToMap(env []string) map[string]string {
	newEnv := make(map[string]string)

	for _, envLine := range env {
		kv := strings.Split(envLine, "=")
		newEnv[kv[0]] = kv[1]
	}

	return newEnv
}

func (s *Details) mapToEnv(env map[string]string) []string {
	newEnv := make([]string, 0)

	for name, val := range env {
		line := fmt.Sprintf("%s=%s", name, val)
		newEnv = append(newEnv, line)
	}

	return newEnv
}
