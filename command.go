package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

// ExecuteCommand executes a shell command with arguments and returns the output (or an error)
func ExecuteCommand(command string, timeout int) (string, error) {
	var shell string
	var args []string
	var timeout_sec = time.Duration(timeout) * time.Second

	// Detect the operating system
	if runtime.GOOS == "windows" {
		// Use cmd.exe on Windows
		shell = "cmd"
		args = []string{"/C", command}
	} else {
		// Use bash (or sh) on Unix-like systems
		shell = "/bin/bash"
		args = []string{"-c", command}
	}

	cmd := exec.CommandContext(context.Background(), shell, args...)
	if timeout_sec > 0 {
		var cancel context.CancelFunc
		cmdContext, cancel := context.WithTimeout(context.Background(), timeout_sec)
		defer cancel()
		cmd = exec.CommandContext(cmdContext, shell, args...)
	}

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	// Handle errors
	if err != nil {
		return fmt.Sprintf("Error: %s\nStderr: %s", err.Error(), stderr.String()), err
	}

	return out.String(), nil
}
