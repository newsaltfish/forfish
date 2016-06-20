package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 md5加密
//  返回md5加密后的字符串
func Md5(origin string) string {
	m := md5.New()
	m.Write([]byte(origin))
	return hex.EncodeToString(m.Sum(nil))
}
