package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// GetSHA256String 将字符串进行SHA256加密并返回结果
func GetSHA256String(s string) string {
	h := sha256.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}
