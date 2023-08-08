package utils

import "fmt"

func Debugger(tag string, format string, args ...interface{}) {
	data := fmt.Sprintf(format, args...)
	if tag == "" {
		tag = "logger"
	}
	fmt.Printf("\n=========== debug start ===========\n%s: %s\n=========== debug end ===========\n", tag, data)

}
