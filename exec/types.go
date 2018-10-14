package exec

import "github.com/mexisme/multiconfig/env"

// Envs interfaces to a type supporting the Combines() method
type Envs interface {
	ToOsEnviron() (env.Envs, error)
}

// Details is simply the struct method-wrapper for the "exec" package
type Details struct {
	envs    Envs
	command []string
}

// New creates a new exec.Details struct
func New() *Details {
	return &Details{}
}

// WithEnvs creates a new exec.Details struct with the .env KV map object copied-in
func (s *Details) WithEnvs(envs Envs) *Details {
	clone := *s // This does a shallow clone

	clone.envs = envs

	return &clone
}

// WithCommand creates a new exec.Details struct with the command []string copied-in
func (s *Details) WithCommand(command []string) *Details {
	clone := *s // This does a shallow clone

	clone.command = command

	return &clone
}
