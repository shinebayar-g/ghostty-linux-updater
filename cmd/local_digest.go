package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func calculateLocalDigest(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Info("file doesn't exist", "file", filePath, "err", err)
			return ""
		} else {
			logger.Error("couldn't open the file", "file", filePath, "err", err)
			os.Exit(1)
		}
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Error("couldn't close the file", "file", filePath, "err", err)
			os.Exit(1)
		}
	}()

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		logger.Error("couldn't calculate the digest of the file", "file", filePath, "err", err)
		os.Exit(1)
	}
	return fmt.Sprintf("sha256:%x", h.Sum(nil))
}
