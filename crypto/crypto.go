package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Md5(in string) string {
	hash := md5.New()
	hash.Write([]byte(in))
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha256(in string) string {
	hash := sha256.New()
	hash.Write([]byte(in))
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha512(in string) string {
	hash := sha256.New()
	hash.Write([]byte(in))
	return hex.EncodeToString(hash.Sum(nil))
}
