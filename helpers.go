package objx

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashWithKey(data, key string) string {
	hash := sha1.New()
	hash.Write([]byte(data + ":" + key))
	return hex.EncodeToString(hash.Sum(nil))
}
