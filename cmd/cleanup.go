package main

import "os"

func cleanup(outDir, oldFilename, newFilename, sigFilename string) {
	logger.Info("Removing the old version", "file", oldFilename)
	err := os.Remove(oldFilename)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Info("file doesn't exist", "file", oldFilename, "err", err)
		} else {
			logger.Error("couldn't remove the file", "file", oldFilename, "err", err)
			os.Exit(1)
		}
	}

	logger.Info("Renaming file", "from", newFilename, "to", oldFilename)
	err = os.Rename(newFilename, oldFilename)
	if err != nil {
		logger.Error("couldn't rename the file.", "file", newFilename, "err", err)
		os.Exit(1)
	}

	logger.Info("Removing the directory", "dir", outDir)
	err = os.RemoveAll(outDir)
	if err != nil {
		logger.Error("couldn't remove the directory", "dir", outDir, "err", err)
		os.Exit(1)
	}

	logger.Info("Removing the minisig file", "file", sigFilename)
	err = os.Remove(sigFilename)
	if err != nil {
		logger.Error("couldn't remove the minisig file.", "file", sigFilename, "err", err)
		os.Exit(1)
	}
}
