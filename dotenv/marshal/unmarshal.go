package marshal

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/subosito/gotenv"
	// "github.com/hashicorp/hcl"
	// "github.com/magiconair/properties"
	toml "github.com/pelletier/go-toml"
	// "github.com/spf13/cast"
	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

func (s *BodyEnv) mapIntoEnv(tree interface{}) map[string]string {
	env := make(map[string]string)

	switch tree.(type) {
	case gotenv.Env:
		for k, v := range tree.(gotenv.Env) {
			env[k] = v
		}
	case map[string]interface{}:
		for k, v := range tree.(map[string]interface{}) {
			env[k] = fmt.Sprintf("%v", v)
		}
	}

	return env
}

func (s *BodyEnv) unmarshallToEnv() error {
	extn := strings.ToLower(filepath.Ext(s.path))

	switch extn {
	case ".env":
		// gotenv.Env == map[string]string, so don't need to convert it
		s.env = s.mapIntoEnv(gotenv.Parse(strings.NewReader(s.body)))

	case ".toml":
		tree, err := toml.Load(s.body)
		if err != nil {
			return ParseError{err}
		}

		s.env = s.mapIntoEnv(tree.ToMap())

	case ".yaml", ".yml":
		if err := yaml.Unmarshal([]byte(s.body), &s.env); err != nil {
			return ParseError{err}
		}

	case ".json":
		var tree map[string]interface{}
		if err := json.Unmarshal([]byte(s.body), &tree); err != nil {
			return ParseError{err}
		}

		s.env = s.mapIntoEnv(tree)

	case ".properties", ".props", ".prop":
		return ParseError{fmt.Errorf(".properties format is not yet supported")}
	}

	log.WithFields(log.Fields{"Env": s.env}).Debug("Unmarshalled!")
	return nil
}
