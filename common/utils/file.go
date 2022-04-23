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
