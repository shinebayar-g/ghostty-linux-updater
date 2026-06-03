package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func untar(filePath string) (dirname string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Error("couldn't close the file", "err", err)
		}
	}()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := gzr.Close(); err != nil {
			logger.Error("couldn't close the gzip reader", "err", err)
		}
	}()

	tr := tar.NewReader(gzr)

	var outputDir string

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, 0755); err != nil {
				return "", err
			}
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				return "", err
			}
			defer func() {
				if err := outFile.Close(); err != nil {
					logger.Error("couldn't close the file", "err", err)
				}
			}()

			if _, err := io.Copy(outFile, tr); err != nil {
				return "", err
			}
		case tar.TypeXGlobalHeader:
			continue
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, header.Name); err != nil {
				return "", err
			}
		default:
			return "", fmt.Errorf("unsupported file type: %c in %s", header.Typeflag, header.Name)
		}

		name := strings.TrimPrefix(header.Name, "./")
		parts := strings.SplitN(name, "/", 2)
		if len(parts) > 0 && parts[0] != "" {
			outputDir = parts[0]
		}
	}

	return outputDir, nil
}
