package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

/*
 1. Fetch the tip release information (digest and download_url).
 2. Calculate the "ghostty-source-current.tar.gz" file's digest if exists.
 3. Compare the digests. If they are the same, Ghostty is up to date!
    If not, then download the tip release's assets.
 4. Verify the "ghostty-source.tar.gz" file's signature using
    the "ghostty-source.tar.gz.minisig" file.
 5. If signature looks good, extract the tarball.
 6. Run "zig build -p $HOME/.local -Doptimize=ReleaseFast"
 7. Remove the previous "ghostty-source-current.tar.gz" file.
 8. Rename the new "ghostty-source.tar.gz" to "ghostty-source-current.tar.gz".
*/
func main() {
	newFilename := "ghostty-source.tar.gz"
	sigFilename := newFilename + ".minisig"
	oldFilename := "ghostty-source-current.tar.gz"
	ghosttyTarball := fetchTipRelease()
	localDigest := calculateLocalDigest(oldFilename)
	if localDigest == ghosttyTarball.source.Digest {
		logger.Info("Digests match. Ghostty is up to date!")
		fmt.Print(ghosttyVersion())
		os.Exit(0)
	} else {
		logger.Info("Digests mismatch. Downloading the latest version...")
		downloadAsset(ghosttyTarball.source)
		downloadAsset(ghosttyTarball.minisig)
		logger.Info("Download complete!")
	}
	checkMiniSign(newFilename)

	logger.Info("Extracting the file", "file", newFilename)
	outDir, err := untar(newFilename)
	if err != nil {
		logger.Error("couldn't extract the file", "file", newFilename, "err", err)
		os.Exit(1)
	}

	logger.Info("Building and installing Ghostty")
	tipVersion := strings.TrimPrefix(outDir, "ghostty-")
	zigBuildOutDir := filepath.Join(os.Getenv("HOME"), ".local")
	cmd := exec.Command("zig", "build", "-p", zigBuildOutDir, "-Doptimize=ReleaseFast", "-Dversion-string="+tipVersion)
	cmd.Dir = outDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		logger.Error("couldn't build Ghostty", "err", err)
		os.Exit(1)
	}

	cleanup(outDir, oldFilename, newFilename, sigFilename)
	logger.Info("Ghossty is updated successfully!")
	fmt.Print(ghosttyVersion())
}
