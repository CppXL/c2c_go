package utils

import (
	"errors"
	"fmt"
)

func IsNumInRange(a, low, high int64, mode [2]byte) (bool, error) {
	// 判断模式是否正确
	for _, mod := range mode {
		if mod != '(' && mod != ')' && mod != '[' && mod != ']' {
			return false, fmt.Errorf("mode value error %c%c", mode[0], mode[1])
		}
	}
	// 判断区间
	if low > high {
		return false, errors.New("low bigger than high")
	}
	// 根据模式进行判断
	switch mode[0] {
	case '[':
		switch mode[1] {

		case ']':
			if a >= low && a <= high {
				return true, nil
			} else {
				return false, nil
			}
		case ')':
			if a >= low && a < high {
				return true, nil
			} else {
				return false, nil
			}
		default:
			return false, fmt.Errorf("mode[0] value error, got %v", mode[0])
		}
	case '(':
		switch mode[1] {
		case ']':
			if a > low && a <= high {
				return true, nil
			} else {
				return false, nil
			}
		case ')':
			if a > low && a < high {
				return true, nil
			} else {
				return false, nil
			}
		default:
			return false, fmt.Errorf("mode[1] value error, got %v", mode[0])
		}
	default:
		return false, fmt.Errorf("mode[0] value error, got %v", mode[0])
	}
}
