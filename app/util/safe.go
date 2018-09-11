package util

import (
	"crypto/md5"
	"fmt"
)

func PasswordEncrypt(password string) string {
	has := md5.Sum([]byte(password))
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}