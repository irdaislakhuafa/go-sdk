package cryptography

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

type SHA256[T any] interface {
	WithKey(key []byte) *T
	Build(ctx context.Context) (string, error)
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

// Generate SHA-256 string value from `text` and return error with code `codes.CodeInvalidValue` if generate is failed
func (s *sha256impl) Build(ctx context.Context) (string, error) {
	if s.key == nil {
		hash := sha256.New()
		if _, err := hash.Write(s.text); err != nil {
			return "", errors.NewWithCode(codes.CodeInvalidValue, "failed to write value to hash, %v", err)
		}
		return hex.EncodeToString(hash.Sum(nil)), nil
	} else {
		hash := hmac.New(sha256.New, s.key)
		if _, err := hash.Write(s.text); err != nil {
			return "", errors.NewWithCode(codes.CodeInvalidValue, "failed to write value to hash, %v", err)
		}
		return hex.EncodeToString(hash.Sum(nil)), nil
	}
}

// Added key to hash SHA-256
func (s *sha256impl) WithKey(key []byte) *sha256impl {
	s.key = key
	return s
}
