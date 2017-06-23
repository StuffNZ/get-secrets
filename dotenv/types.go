package dotenv

import (
	"github.com/subosito/gotenv"
)

// BodyEnv TODO
type BodyEnv struct {
	path string
	body string
	env  gotenv.Env
}

// BodyEnvMap TODO
type BodyEnvMap map[string]BodyEnv

// DotEnvs TODO
type DotEnvs struct {
	env BodyEnvMap
}

// New TODO
func New() *DotEnvs {
	return &DotEnvs{env: make(BodyEnvMap)}
}
