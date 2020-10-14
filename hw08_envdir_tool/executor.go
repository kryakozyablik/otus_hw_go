package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	cmdName := cmd[0]
	cmdArgs := cmd[1:]

	execCmd := exec.Command(cmdName, cmdArgs...)
	execCmd.Stdin = os.Stdin
	execCmd.Stderr = os.Stderr
	execCmd.Stdout = os.Stdout

	for k, v := range env {
		if len(v) == 0 {
			_ = os.Unsetenv(k)
		} else {
			_ = os.Setenv(k, v)
		}
	}

	if err := execCmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			returnCode = exitError.ExitCode()
		}
	}

	return
}
