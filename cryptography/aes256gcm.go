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

	result := []byte{}
	gcmSeal := gcm.Seal(nonce, nonce, text, nil)
	_ = hex.Encode(result, gcmSeal)
	base64.StdEncoding.Encode(result, result)
	return result, nil
}

// func DecryptAES256GCM(text, key []byte) ([]byte, error)
