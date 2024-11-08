package util

import "fmt"

func Float64ToString2f(val float64) string {
	return fmt.Sprintf("%.2f", val)
}

func InterfaceToString2f(val interface{}) string {
	return fmt.Sprintf("%.2f", val.(float64))
}
