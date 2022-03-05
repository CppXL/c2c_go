package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetCurrDir() string {
	return getCurrDir()
}

func getCurrDir() string {
	path, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	dir := filepath.Dir(path)
	return dir + "/"
}
