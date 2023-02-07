package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetCurrDir(t *testing.T) {
	fmt.Println(GetCurrentDir())
	fmt.Println(os.Executable())
	fmt.Println(filepath.Dir(GetCurrentDir() + "../../srver"))

	fmt.Println(filepath.Dir(filepath.Dir(GetCurrentDir())))
	var a int
	fmt.Scan(&a)
}
