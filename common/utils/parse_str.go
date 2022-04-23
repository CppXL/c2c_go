package utils

// 去除字符串之前和字符串之后的空格
func RemoveSpace(str string) string {
	for i := range str {
		if i > 0 && str[i] == ' ' && str[i-1] == ' ' {
			str = str[:i] + str[i+1:]
		}
	}
	return ""
}
