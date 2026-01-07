package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
)

func loadJSONFromFile(path string, target interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", path, err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal JSON from file %s: %w", path, err)
	}

	return nil
}

func createDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}

	return nil
}

func getAbsolutePath(relativePath string) (string, error) {
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for %s: %w", relativePath, err)
	}

	return absPath, nil
}

func handleMultiError(err error) {
	if err != nil {
		var merr *multierror.Error
		if errors.As(err, &merr) {
			for _, e := range merr.Errors {
				log.Printf("error: %v", e)
			}
		} else {
			log.Printf("error: %v", err)
		}
	}
}