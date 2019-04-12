package util

import "strings"

func Trim(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}