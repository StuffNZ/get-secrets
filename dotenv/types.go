package dotenv

import (
	"github.com/subosito/gotenv"
)

// BodyEnv contains the given "path" and "body" of a .env, as well as the parsed .env KV map
type BodyEnv struct {
	path string
	body string
	env  gotenv.Env
}

// BodyEnvMap contains map of "path" --> BodyEnv structs
type BodyEnvMap map[string]BodyEnv

// DotEnvs is simply the struct method-wrapper for the "dotenv" package
type DotEnvs struct {
	env BodyEnvMap
}

// New creates a new DotEnvs struct
func New() *DotEnvs {
	return &DotEnvs{env: make(BodyEnvMap)}
}
