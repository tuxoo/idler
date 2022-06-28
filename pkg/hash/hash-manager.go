package hash

import (
	"crypto/sha1"
	"fmt"
	"hash"
)

type SHA1Hasher struct {
	salt   string
	hasher hash.Hash
}

func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{
		salt:   salt,
		hasher: sha1.New(),
	}
}

func (h *SHA1Hasher) Hash(content string) string {
	h.hasher.Write([]byte(content))

	return fmt.Sprintf("%x", h.hasher.Sum([]byte(h.salt)))
}
