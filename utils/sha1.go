package utils

import (
	"crypto/sha1"
	"encoding/base64"
)

func Hash(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	sha1 := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha1
}
