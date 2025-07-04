package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(txt string) string {
	md := md5.New()
	byt := md.Sum([]byte(txt))
	return hex.EncodeToString(byt)
}
