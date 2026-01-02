package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func untar(file string) (dirname string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return "", err
	}
	defer gzr.Close()

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
			defer outFile.Close()

			if _, err := io.Copy(outFile, tr); err != nil {
				return "", err
			}
		case tar.TypeXGlobalHeader:
			continue
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
