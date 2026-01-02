package main

import (
	"fmt"
	"os"
	"os/exec"
)

func ghosttyVersion() string {
	cmd := exec.Command("ghostty", "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("couldn't check the Ghossty version", "err", err)
		os.Exit(1)
	}
	return fmt.Sprintf("%s", out)
}
