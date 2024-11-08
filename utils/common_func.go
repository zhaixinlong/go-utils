package util

import "strings"

// 是否包含-适用于string
func IsContainForString(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
func IsContain(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 反转
func ReverseStringWithComma(str string) string {
	splitStr := strings.Split(str, ",")
	for i, j := 0, len(splitStr)-1; i < j; i, j = i+1, j-1 {
		splitStr[i], splitStr[j] = splitStr[j], splitStr[i]
	}
	return strings.Join(splitStr, ",")
}
