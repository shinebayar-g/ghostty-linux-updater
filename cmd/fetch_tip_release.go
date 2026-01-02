package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
)

type Asset struct {
	Name     string `json:"name"`
	Uploader struct {
		Login string `json:"login"`
		ID    int    `json:"id"`
	} `json:"uploader"`
	Digest string `json:"digest"`
	URL    string `json:"browser_download_url"`
}

type fetchTipReleaseResponse struct {
	Assets []Asset `json:"assets"`
}

type GhosttyTarball struct {
	source  Asset
	minisig Asset
}

func checkUploader(a Asset) bool {
	logger.Info("Checking the uploader...")
	if a.Uploader.Login == "mitchellh" && a.Uploader.ID == 1299 {
		logger.Info("Uploader is verified.")
		return true
	} else {
		logger.Info("Uploader is not verified!")
		logger.Error("uploader login information", "login", a.Uploader.Login, "id", a.Uploader.ID)
		return false
	}
}

func fetchTipRelease() GhosttyTarball {
	ctx := context.Background()
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com/repos/ghostty-org/ghostty/releases/tags/tip", nil)
	if err != nil {
		logger.Error("couldn't create the request", "err", err)
		os.Exit(1)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("couldn't fetch the tip release", "err", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var response fetchTipReleaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		logger.Error("couldn't decode the response body", "err", err)
		os.Exit(1)
	}

	var ghosttyTarball GhosttyTarball
	for _, asset := range response.Assets {
		switch asset.Name {
		case "ghostty-source.tar.gz":
			logger.Info("Found the ghostty-source.tar.gz from the release assets")
			if !checkUploader(asset) {
				os.Exit(1)
			}
			ghosttyTarball.source = asset
		case "ghostty-source.tar.gz.minisig":
			logger.Info("Found the ghostty-source.tar.gz.minisig from the release assets")
			if !checkUploader(asset) {
				os.Exit(1)
			}
			ghosttyTarball.minisig = asset
		}
	}

	return ghosttyTarball
}
