package validators

import (
	"log"
	"os"
	"path/filepath"
)

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm) // 0755 by default
	if err != nil {
		return err
	}
	return nil
}

func setupDirectories() {
	baseDir := filepath.Join("media", "inumazioni", "foto")
	if err := ensureDir(baseDir); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}
}
