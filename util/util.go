package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Trim(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}

func GetMd5(str string) string {
	t := md5.Sum([]byte(str))
	return hex.EncodeToString(t[:])
}