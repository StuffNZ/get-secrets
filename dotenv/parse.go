package dotenv

import (
	"fmt"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

/*
AddFromString allows to add new *.env string to the list of DotEnvs

- Need to provide the body-content as a string.

- Need to provide a "path" to where the .env was read from, to allow for combination based on lexical-order of this path.

- This gets parsed immediately, not deferred until combined.
*/
// TODO: Is "path" the wrong name for this?
func (s *DotEnvs) AddFromString(path string, body string) error {
	if path == "" {
		return fmt.Errorf("Empty path provided")
	}
	if body == "" {
		return fmt.Errorf("Empty body from %#v", path)
	}

	log.WithFields(log.Fields{"path": path, "body": body}).Debug("Parsing dotenv file...")
	s.env[path] = BodyEnv{path: path, body: body, env: s.parseEnv(body)}

	return nil
}

/*
Combine Combine all the .env's provided into a single .env map

Returns a new .env map (gotenv.Env) with the combined .env's.
The .env's are combined by lexical-ordering of the "path" field;

i.e. an .env with a "path" of "z..." will be added after those with a "path" of "a..."
*/
// TODO: Need to make sure file's named ".env" (only) are parsed last, somehow...
func (s *DotEnvs) Combine() map[string]string {
	joinedEnv := make(gotenv.Env)

	for _, path := range s.sortedPaths() {
		s.mergeEnv(s.env[path].env, joinedEnv)
	}

	return joinedEnv
}

func (s *DotEnvs) mergeEnv(from, to map[string]string) {
	for k, v := range from {
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
