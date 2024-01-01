package cryptography

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"golang.org/x/crypto/argon2"
)

type Argon2 interface {
	// Hash text with argon2 algorithm with format from `hashFormat` or you can determine your specific format
	Hash(text []byte) (string, error)

	// Compare plain text with hash argon2, return `true`` and error `nil` if equal
	Compare(text, hashedText []byte) (bool, error)
}

type argon2impl struct {
	delimiter   string
	lengthValue uint64
	hashFormat  string
	salt        []byte
	iterations  uint32
	parallelism uint8
	keyLen      uint32
	memory      uint32
	version     int
}

// Create argon2 hash with default parameter
func NewArgon2() Argon2 {
	result := &argon2impl{
		delimiter:   "$",
		lengthValue: 6,
		hashFormat:  "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		salt:        make([]byte, 16),
		iterations:  3,
		parallelism: 1,
		keyLen:      32,
		memory:      (4 * 1024),
		version:     argon2.Version,
	}

	return result
}

func (a *argon2impl) Hash(text []byte) (string, error) {
	if _, err := rand.Read(a.salt); err != nil {
		return "", errors.NewWithCode(codes.CodeArgon2, "failed to make salt, %v", err)
	}

	key := argon2.IDKey(text, a.salt, a.iterations, a.memory, a.parallelism, a.keyLen)
	encodedKey, err := a.encodeKey(key)
	if err != nil {
		return "", errors.NewWithCode(codes.CodeArgon2, "failed to encode key with argon2, %v", err)
	}

	return encodedKey, nil
}

func (a *argon2impl) Compare(text []byte, hashedText []byte) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (a *argon2impl) encodeKey(key []byte) (string, error) {
	enc := base64.StdEncoding
	base64Salt := enc.EncodeToString(a.salt)
	base64Key := enc.EncodeToString(key)
	encodedKey := fmt.Sprintf(a.hashFormat, a.version, a.memory, a.iterations, a.parallelism, base64Salt, base64Key)
	return encodedKey, nil
}
