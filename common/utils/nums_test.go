package utils

import (
	"fmt"
	"testing"
)

func TestIsNumInRange(t *testing.T) {
	numSet := []struct {
		a    int64
		low  int64
		high int64
		mode [2]byte
		want bool
	}{
		{int64(1), int64(1), int64(2), [2]byte{'[', ']'}, true},
		{int64(2), int64(1), int64(2), [2]byte{'[', ']'}, true},
		{int64(3), int64(1), int64(3), [2]byte{'(', ']'}, true},
		{int64(1), int64(1), int64(3), [2]byte{'(', ']'}, false},
		{int64(1), int64(1), int64(3), [2]byte{'[', ')'}, true},
		{int64(3), int64(1), int64(3), [2]byte{'[', ')'}, false},
		{int64(1), int64(1), int64(3), [2]byte{'(', ')'}, false},
		{int64(3), int64(1), int64(3), [2]byte{'(', ')'}, false},

		{int64(10), int64(1), int64(56), [2]byte{'[', ')'}, true},
		{int64(12), int64(13), int64(53), [2]byte{'[', ']'}, false},
		{int64(4), int64(1), int64(3), [2]byte{'[', ']'}, false},
		{int64(112), int64(1), int64(112), [2]byte{'(', ')'}, false},
		{int64(111), int64(1), int64(1), [2]byte{'7', 'k'}, false},
		{int64(111), int64(345), int64(1), [2]byte{'[', ']'}, false},
	}

	for _, test := range numSet {
		if out, err := IsNumInRange(test.a, test.low, test.high, test.mode); err == nil && out == test.want {
			continue
		} else if err == nil && out != test.want {
			fmt.Printf("data %v error", test)
		} else if err != nil {
			fmt.Println("error:", err)
		}

	}
}
