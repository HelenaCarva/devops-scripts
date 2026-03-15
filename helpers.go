package devops_scripts

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GenerateRandomString generates a random string of a specified length
func GenerateRandomString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	b := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// HashString generates a SHA-256 hash of a string
func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// GetAbsolutePath returns the absolute path of a file or directory
func GetAbsolutePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	return absPath, nil
}

// CheckFileExists checks if a file exists
func CheckFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// LogError logs an error message
func LogError(err error) {
	log.Printf("Error: %v", err)
}

// TrimString trims whitespace from the beginning and end of a string
func TrimString(s string) string {
	return strings.TrimSpace(s)
}

// CheckDirExists checks if a directory exists
func CheckDirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func main() {
	// Example usage:
	randomString, err := GenerateRandomString(10)
	if err != nil {
		LogError(err)
	} else {
		fmt.Println(randomString)
	}

	hash := HashString("example")
	fmt.Println(hash)

	absPath, err := GetAbsolutePath(".")
	if err != nil {
		LogError(err)
	} else {
		fmt.Println(absPath)
	}

	if CheckFileExists("helpers.go") {
		fmt.Println("File exists")
	} else {
		fmt.Println("File does not exist")
	}

	trimmedString := TrimString("   example   ")
	fmt.Println(trimmedString)

	if CheckDirExists("..") {
		fmt.Println("Directory exists")
	} else {
		fmt.Println("Directory does not exist")
	}
}