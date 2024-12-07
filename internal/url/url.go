package url

import (
	"crypto/sha256"
	"encoding/hex"
)

func Shorten(originalUrl string) string {
	hash := sha256.New()
	hash.Write([]byte(originalUrl))
	hashed := hex.EncodeToString(hash.Sum(nil))
	hashedUrl := hashed[:8]
	return hashedUrl
}
