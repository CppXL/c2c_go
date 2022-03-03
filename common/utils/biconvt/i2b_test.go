package biconv

import (
	"fmt"
	"testing"
)

func TestI32tob(t *testing.T) {
	for i, v := range I32tob(uint32(10)) {
		fmt.Println(i, v)
	}
}

func TestI64tob(t *testing.T) {
	for i, v := range I64tob(uint64(10)) {
		fmt.Println(i, v)
	}
}
