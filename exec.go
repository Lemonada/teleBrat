package main

import (
	"bytes"
	"os/exec"
	"syscall"
	"time"
)

func execute(command string) (string, string, error) {
	newCommand := "(" + command + ")"
	cmd := exec.Command("cmd.exe", "/C", newCommand)
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	cmd.Start()

	// Use a channel to signal completion so we can use a select statement
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	// Start a timer
	timeout := time.After(60 * time.Second)

	select {
	case <-timeout:
		// Timeout happened first, kill the process and print a message.
		cmd.Process.Kill()
		return "Command timed out", " ", nil
	case err := <-done:
		// Command completed before timeout. Print output and error if it exists.
		if err != nil {
			return "", outErr.String(), err
		}
		return out.String(), " ", nil
	}
}
