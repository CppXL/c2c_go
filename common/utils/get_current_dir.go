package utils

import (
	"log"
	"os"
	"path/filepath"
)

// Get executable file path
// return string
func GetCurrentDir() string {
	path, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	dir := filepath.Dir(path)
	return dir + "/"
}
