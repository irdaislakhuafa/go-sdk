package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

// Encrypt data with AES-256-GCM algorithm
func EncryptAES256GCM(text, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, "failed to create new cipher block, %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, "failed to create new ciper gcm, %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, "failed to read random value for nonce, %v", err)
	}

	s := string(gcm.Seal(nonce, nonce, text, nil))
	s = hex.EncodeToString([]byte(s))
	s = base64.StdEncoding.EncodeToString([]byte(s))

	return []byte(s), nil
}

// Decrypt data with AES-256-GCM algorithm
// func DecryptAES256GCM(text, key []byte) ([]byte, error)
