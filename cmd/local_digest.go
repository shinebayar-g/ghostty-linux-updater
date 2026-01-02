package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func calculateLocalDigest(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Info("file doesn't exist", "file", filename, "err", err)
			return ""
		} else {
			logger.Error("couldn't open the file", "file", filename, "err", err)
			os.Exit(1)
		}
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		logger.Error("couldn't calculate the digest of the file", "file", filename, "err", err)
		os.Exit(1)
	}
	return fmt.Sprintf("sha256:%x", h.Sum(nil))
}
