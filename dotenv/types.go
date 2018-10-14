package dotenv

import (
	"bitbucket.org/mexisme/get-secrets/dotenv/env"
)

// BodyEnvMap contains map of "path" --> BodyEnv structs
type BodyEnvMap map[string]*env.BodyEnv

// DotEnvs is simply the struct method-wrapper for the "dotenv" package
type DotEnvs struct {
	env BodyEnvMap
}

// New creates a new DotEnvs struct
func New() *DotEnvs {
	return &DotEnvs{env: make(BodyEnvMap)}
}
