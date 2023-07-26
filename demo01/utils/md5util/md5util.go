package md5util

import (
	"crypto/md5"
	"fmt"
)

// md5加密方法
func MD5(str string) string {
	data := []byte(str)
	sum := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", sum)
	return md5Str
}

func CreateSecret(password string) string {
	pass := MD5(MD5(password))
	return pass
}
