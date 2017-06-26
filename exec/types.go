package exec

// DotEnv interfaces to a type supporting the Combins() method
type DotEnv interface {
	Combine() map[string]string
}

// Details is simply the struct method-wrapper for the "exec" package
type Details struct {
	env     []string
	dotEnvs DotEnv
	command []string
}

// New creates a new exec.Details struct
func New() *Details {
	return &Details{}
}

// WithDotEnvs creates a new exec.Details struct with the .env KV map object copied-in
func (s *Details) WithDotEnvs(dotEnvs DotEnv) *Details {
	clone := *s // This does a shallow clone

	clone.dotEnvs = dotEnvs

	return &clone
}

// WithEnviron creates a new exec.Details struct with the os.Environ object copied-in
func (s *Details) WithEnviron(env []string) *Details {
	clone := *s // This does a shallow clone

	clone.env = env

	return &clone
}

// WithCommand creates a new exec.Details struct with the command []string copied-in
func (s *Details) WithCommand(command []string) *Details {
	clone := *s // This does a shallow clone

	clone.command = command

	return &clone
}
