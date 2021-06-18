package util

import (
	"crypto/md5"
	"fmt"
)

// GetMd5 获取 md5 值
func GetMd5(strData string) string {
	if strData == "" {
		return strData
	}

	return fmt.Sprintf("%x", md5.Sum([]byte(strData)))
}
