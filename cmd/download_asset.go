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
	defer resp.Body.Close()

	file, err := os.Create(asset.Name)
	if err != nil {
		logger.Error("couldn't create the file", "err", err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		logger.Error("couldn't write the file", "err", err)
		os.Exit(1)
	}
}
