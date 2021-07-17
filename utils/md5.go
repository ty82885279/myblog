package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "ty11223344"

func EncryptPassword(psw string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(psw)))
}
