package env

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// BodyEnv contains the given "path" and "body" of a .env, as well as the parsed .env KV map
type BodyEnv struct {
	path string
	body string
	env  map[string]string
}

// ParseError when failing to parse an env file
type ParseError struct {
	err error
}

// New object
func New() *BodyEnv {
	return &BodyEnv{}
}

// WithPathBody creates a new BodyEnv with Path and Body added, and the Env updated
func (s *BodyEnv) WithPathBody(path, body string) *BodyEnv {
	clone := *s // Shallow clone
	clone.path = path
	clone.body = body
	// Explicitly reset it:
	clone.env = nil
	clone.Env()

	return &clone
}

// Env returns the env map
func (s *BodyEnv) Env() map[string]string {
	if s.env == nil {
		s.PanicOnBadAttributes()
		if err := s.unmarshallToEnv(); err != nil {
			log.WithFields(log.Fields{"Error": err}).Panicf("Failed to parse '%v'", s.path)
		}
	}
	return s.env
}

// PanicOnBadAttributes throws a panic if any of the necessary attrs are missing
func (s *BodyEnv) PanicOnBadAttributes() {
	if s.path == "" {
		log.Panic("Path is empty")
	}
	if s.body == "" {
		log.Panic("Body is empty")
	}
}

// Error returns the formatted configuration error.
func (s ParseError) Error() string {
	return fmt.Sprintf("While parsing config: %s", s.err.Error())
}
