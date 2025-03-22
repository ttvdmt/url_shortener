package encode

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateCode(length uint64) string {
	res := make([]byte, length)
	_, _ = rand.Read(res)

	return base64.URLEncoding.EncodeToString(res)[:length]
}
