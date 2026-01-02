package main

import (
	"fmt"
	"os"
	"os/exec"
)

func checkMiniSign(filename string) {
	logger.Info("Checking minisign signature")
	// Key is published here
	// https://ghostty.org/docs/install/build#getting-the-source-code
	minisignKey := "RWQlAjJC23149WL2sEpT/l0QKy7hMIFhYdQOFy0Z7z7PbneUgvlsnYcV"
	cmd := exec.Command("minisign", "-Vm", filename, "-P", minisignKey)
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s", out)
	if err != nil {
		logger.Error("couldn't verify the minisign signature", "err", err)
		os.Exit(1)
	}
	logger.Info("file signature match")
}
