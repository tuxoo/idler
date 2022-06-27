package hash

import (
	"crypto/sha1"
	"fmt"
)

// SHA1Hasher uses SHA1 to hash passwords with provided salt
type SHA1Hasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{salt: salt}
}

func (h *SHA1Hasher) Hash(content string) string {
	hash := sha1.New()
	hash.Write([]byte(content))

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt)))
}
