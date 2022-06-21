package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

const (
	CodeOK   = 0
	CodeFail = 1
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(args []string, env Environment) int {
	if len(args) == 0 {
		log.Println("you need to specify the program to run")
		return CodeFail
	}

	for name, value := range env {
		var err error
		if value.NeedRemove {
			err = os.Unsetenv(name)
		} else {
			err = os.Setenv(name, value.Value)
		}

		if err != nil {
			return CodeFail
		}
	}

	//nolint:gosec
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	err := cmd.Run()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}

		return CodeFail
	}

	return CodeOK
}
