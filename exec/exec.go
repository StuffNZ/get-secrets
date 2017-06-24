package exec

import (
	"os/exec"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// Exec TODO
func (s *Details) Exec() error {
	log.WithFields(log.Fields{"command": s.command[0]}).Debug("Finding FQ Path")
	binPath, err := exec.LookPath(s.command[0])
	if err != nil {
		return err
	}

	env := s.CombineEnvs()
	log.WithFields(log.Fields{"binPath": binPath, "command": s.command, "env": env}).Debug("Running command...")
	return syscall.Exec(binPath, s.command, env)
}
