package util

import (
	"crypto/sha1"
	"encoding/base64"
)

// Hash ...
func Hash(str string) (hash string) {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	hash = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return hash
}
