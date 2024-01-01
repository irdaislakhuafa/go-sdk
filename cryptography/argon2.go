package cryptography

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

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
	b64enc      base64.Encoding
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
		b64enc:      *base64.StdEncoding.Strict(),
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
	// extract all parameter include salt and key length from encoded argon2 hash
	arg, key, err := a.decodeKey(hashedText)
	if err != nil {
		return false, err
	}

	// generate other hash with same parameters and compare it with existsing hash
	otherKey := argon2.IDKey(text, arg.salt, arg.iterations, arg.memory, arg.parallelism, arg.keyLen)
	if subtle.ConstantTimeCompare(key, otherKey) != 1 {
		return false, errors.NewWithCode(codes.CodeArgon2NotMatch, "argon2 for plain text with hashed text is not match")
	}

	return true, nil
}

func (a *argon2impl) encodeKey(key []byte) (string, error) {
	base64Salt := a.b64enc.EncodeToString(a.salt)
	base64Key := a.b64enc.EncodeToString(key)
	encodedKey := fmt.Sprintf(a.hashFormat, a.version, a.memory, a.iterations, a.parallelism, base64Salt, base64Key)
	return encodedKey, nil
}

func (a *argon2impl) decodeKey(key []byte) (*argon2impl, []byte, error) {
	// split enncoded argon2 hash with delimiter
	values := strings.Split(string(key), a.delimiter)

	// compare length values with standard length
	if lengthValue := len(values); lengthValue != int(a.lengthValue) {
		return nil, nil, errors.NewWithCode(codes.CodeArgon2, "invalid length of encoded hash, expected %v but got %v", a.lengthValue, lengthValue)
	}

	// check argon2 version compatibility
	version := 0
	if _, err := fmt.Sscanf(values[2], "v=%v", &version); err != nil {
		return nil, nil, errors.NewWithCode(codes.CodeArgon2, "failed to get argon2 version from hash text, %v", err)
	}

	if a.version != version {
		return nil, nil, errors.NewWithCode(codes.CodeArgon2IncompatibleVersion, "current argon2 version is %v but encoded hash using version %v", a.version, version)
	}

	// mapping values for for memory, iterations and parallelism
	arg := &argon2impl{}
	if _, err := fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &arg.memory, &arg.iterations, &arg.parallelism); err != nil {
		return nil, nil, errors.NewWithCode(codes.CodeArgon2, "failed to get memory, iterations and parallelism information from hash, %v", err)
	}

	// decode base64 salt
	var err error
	if arg.salt, err = a.b64enc.DecodeString(values[4]); err != nil {
		return nil, nil, errors.NewWithCode(codes.CodeArgon2, "failed to decode salt information from hash, %v", err)
	}

	// decode base64 key
	key, err = a.b64enc.DecodeString(values[5])
	if err != nil {
		return nil, nil, errors.NewWithCode(codes.CodeArgon2, "failed to decode key information from hash, %v", err)
	}
	arg.keyLen = uint32(len(key))

	return arg, key, nil
}
