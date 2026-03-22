package main

import (
	"io"
	"net/http"
	"os"
)

func downloadAsset(asset Asset) {
	resp, err := http.Get(asset.URL)
	if err != nil {
		logger.Error("couldn't download the asset", "err", err)
		os.Exit(1)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.Error("couldn't close the response body", "err", err)
			os.Exit(1)
		}
	}()

	file, err := os.Create(asset.Name)
	if err != nil {
		logger.Error("couldn't create the file", "err", err)
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Error("couldn't close the file", "err", err)
			os.Exit(1)
		}
	}()

	if _, err := io.Copy(file, resp.Body); err != nil {
		logger.Error("couldn't write the file", "err", err)
		os.Exit(1)
	}
}
