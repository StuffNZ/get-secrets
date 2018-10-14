package exec

import (
	"os/exec"
	"syscall"

	log "github.com/sirupsen/logrus"
)

/*
Exec runs the provided "command".

It looks-up the FQ Path to the command, also combines the given os.Environ (copy) and .env KV maps into a
new os.Environ []string and passes these into the syscall.Exec() function.
*/
func (s *Details) Exec() error {
	binPath, err := s.bin()
	if err != nil {
		return err
	}

	envs, err := s.envs.ToOsEnviron()
	if err != nil {
		return err
	}

	log.Infof("Running command %v (%v)...", s.command, binPath)
	// Need to suppress "warning: Subprocess launched with variable,MEDIUM,HIGH (gosec)"
	return syscall.Exec(binPath, s.command, envs) // nolint: gosec
}

func (s *Details) bin() (string, error) {
	return exec.LookPath(s.command[0])
}
