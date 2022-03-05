package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetCurrDir(t *testing.T) {
	fmt.Println(GetCurrDir())
	fmt.Println(os.Executable())
	fmt.Println(filepath.Dir(GetCurrDir() + "../../srver"))

	fmt.Println(filepath.Dir(filepath.Dir(GetCurrDir())))
	var a int
	fmt.Scan(&a)
}
