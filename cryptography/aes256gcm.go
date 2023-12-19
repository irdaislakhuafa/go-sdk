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

type AES256GCM interface {
	// Encrypt data with AES-256-GCM algorithm
	Encrypt(text []byte) ([]byte, error)

	// Decrypt data with AES-256-GCM algorithm
	Decrypt(text []byte) ([]byte, error)
}

type aes256gcm struct {
	key []byte
}

// Create new encrypt/decrypt with AES-256-GCM algoritm
func NewAES256GCM(key []byte) AES256GCM {
	return &aes256gcm{key: key}
}

func (a *aes256gcm) Encrypt(text []byte) ([]byte, error) {
	return EncryptAES256GCM(text, a.key)
}

func (a *aes256gcm) Decrypt(text []byte) ([]byte, error) {
	return DecryptAES256GCM(text, a.key)
}

// Encrypt data with AES-256-GCM algorithm, like `AES256GCM.Encrypt()` but fully functional
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

	result := string(gcm.Seal(nonce, nonce, text, nil))
	result = hex.EncodeToString([]byte(result))
	result = base64.StdEncoding.EncodeToString([]byte(result))

	return []byte(result), nil
}

// Decrypt data with AES-256-GCM algorithm, like `AES256GCM.Decrypt()` but fully functional
func DecryptAES256GCM(text, key []byte) ([]byte, error) {
	text, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, "failed to decode base64 string, %v", err)
	}

	if text, err = hex.DecodeString(string(text)); err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, "failed to decode hex string, %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, "failed to create new block cipher, %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeInvalidValue, "failed to create new gcm, %v", err)
	}

	nonceSize := gcm.NonceSize()
	nonce, encryptedText := text[:nonceSize], text[nonceSize:]

	result, err := gcm.Open(nil, nonce, encryptedText, nil)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeAES256GCMOpenError, "failed to open encrypted text with AES-256-GCM, %v", err)
	}

	return result, nil
}
