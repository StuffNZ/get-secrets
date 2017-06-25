package exec

// DotEnv TODO
type DotEnv interface {
	Combine() map[string]string
}

// Details TODO
type Details struct {
	env     []string
	dotEnvs DotEnv
	command []string
}

// New TODO
func New() *Details {
	return &Details{}
}

// WithDotEnvs TODO
func (s *Details) WithDotEnvs(dotEnvs DotEnv) *Details {
	clone := *s // This does a shallow clone

	clone.dotEnvs = dotEnvs

	return &clone
}

// WithEnviron TODO
func (s *Details) WithEnviron(env []string) *Details {
	clone := *s // This does a shallow clone

	clone.env = env

	return &clone
}

// WithCommand TODO
func (s *Details) WithCommand(command []string) *Details {
	clone := *s // This does a shallow clone

	clone.command = command

	return &clone
}
