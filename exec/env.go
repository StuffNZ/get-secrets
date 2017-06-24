package exec

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

// CombineEnvs TODO
func (s *Details) CombineEnvs() []string {
	newEnvs := s.envToMap(s.env)
	dotEnvs := s.dotEnvs.Combine()
	s.mergeEnv(dotEnvs, newEnvs)

	log.WithFields(log.Fields{"new": newEnvs}).Debug("Combined envs")

	return s.mapToEnv(newEnvs)
}

// Merge the first map into the second
// Because maps are implicit references, updating the second map updates the caller's map
func (s *Details) mergeEnv(from, to map[string]string) {
	log.WithFields(log.Fields{"from": from, "to": to}).Debug("Merging envs")
	for k, v := range from {
		log.WithFields(log.Fields{"key": k, "val": v}).Debug()
		to[k] = v
	}
}

func (s *Details) envToMap(env []string) map[string]string {
	log.WithFields(log.Fields{"envList": env}).Debug("Converting from env list to map")
	newEnv := make(map[string]string)

	for _, envLine := range env {
		kv := strings.Split(envLine, "=")
		newEnv[kv[0]] = kv[1]
	}

	log.WithFields(log.Fields{"envMap": newEnv}).Debug("Converted from env list to map")

	return newEnv
}

func (s *Details) mapToEnv(env map[string]string) []string {
	log.WithFields(log.Fields{"envMap": env}).Debug("Converting from env map to list")
	newEnv := make([]string, 0)

	for name, val := range env {
		line := fmt.Sprintf("%s=%s", name, val)
		newEnv = append(newEnv, line)
	}

	log.WithFields(log.Fields{"envList": newEnv}).Debug("Converted from env map to list")

	return newEnv
}
