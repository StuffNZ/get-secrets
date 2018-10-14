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
	binPath, err := exec.LookPath(s.command[0])
	if err != nil {
		return err
	}

	env := s.CombineEnvs()
	log.WithFields(log.Fields{"binPath": binPath, "command": s.command, "env": env}).Debug("Running command...")
	return syscall.Exec(binPath, s.command, env)
}
