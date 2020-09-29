package main

import (
	"fmt"
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

	execCmd.Env = os.Environ()
	for k, v := range env {
		execCmd.Env = append(execCmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	if err := execCmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			returnCode = exitError.ExitCode()
		}
	}

	return
}
