package biconv

import (
	"fmt"
	"testing"
)

func TestBtoi64(t *testing.T) {
	fmt.Println(Btoi64([]byte{'a', 'b', 'c', 'd', 'e', 'f', 'r', 'r'}))

}

func TestBtoi32(t *testing.T) {
	fmt.Println(Btoi32([]byte{byte(120), byte(120), byte(120), byte(120)}))

}
