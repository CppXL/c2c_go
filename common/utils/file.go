package utils

import (
	"os"
)

func IsFileExists(filePath string) bool {
	Finfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !Finfo.IsDir()
}

func IsDirExists(path string) bool {
	Finfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return Finfo.IsDir()
}
