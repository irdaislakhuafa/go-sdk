package cryptography

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"hash"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/operator"
)

type SHA256[T any] interface {
	// Added key to hash SHA-256
	WithKey(key []byte) *T

	// Generate SHA-256 string value from `text` and return error with code `codes.CodeInvalidValue` if generate is failed
	Build() (string, error)
}

type sha256impl struct {
	key  []byte
	text []byte
}

// Create hash with SHA-256 algoritm
func NewSHA256(text []byte) SHA256[sha256impl] {
	return &sha256impl{
		text: text,
	}
}

func (s *sha256impl) Build() (string, error) {
	hash := operator.Ternary[hash.Hash](s.key == nil, sha256.New(), hmac.New(sha256.New, s.key))
	if _, err := hash.Write(s.text); err != nil {
		return "", errors.NewWithCode(codes.CodeInvalidValue, "failed to write value to hash, %v", err)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (s *sha256impl) WithKey(key []byte) *sha256impl {
	s.key = key
	return s
}
