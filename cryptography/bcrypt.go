package cryptography

import (
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	BcryptDefaultCost = bcrypt.DefaultCost
	BcryptMaxCost     = bcrypt.MaxCost
	BcryptMinCost     = bcrypt.MinCost
)

type Bcrypt interface {
	// Hash plain text with bcrypt algoritm
	Hash(plainText []byte) ([]byte, error)

	// Compare plain text with hashed text with bcrypt algorithm
	Compare(plainText, hashedText []byte) error

	// Set cost of bcrypt algoritm. By default using `cryptography.BcryptDefaultCost`
	SetCost(cost int) Bcrypt

	// Get cost used of this bcrypt algoritm
	GetCost() int

	// Count cost value from hashed text with bcrypt algorithm. This method will not set `cost` of existing bcrypt
	CostFromHash(hashedText []byte) (int, error)
}

type bcryptimpl struct {
	cost int
}

// Create new bcrypt based on `golang.org/x/crypto/bcrypt` with default `cost` using `cryptography.BcryptDefaultCost`
func NewBcrypt() Bcrypt {
	result := &bcryptimpl{
		cost: BcryptDefaultCost,
	}

	return result
}

func (b *bcryptimpl) Hash(plainText []byte) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword(plainText, b.cost)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeBcrypt, "failed to hash plain text with bcrypt, %v", err)
	}

	return result, nil
}

func (b *bcryptimpl) Compare(plainText []byte, hashedText []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedText, plainText); err != nil {
		return errors.NewWithCode(codes.CodeBcrypt, "plain text with hashed text not match, %v", err)
	}
	return nil
}

func (b *bcryptimpl) SetCost(cost int) Bcrypt {
	b.cost = cost
	return b
}

func (b *bcryptimpl) GetCost() int {
	return b.cost
}

func (b *bcryptimpl) CostFromHash(hashedText []byte) (int, error) {
	cost, err := bcrypt.Cost(hashedText)
	if err != nil {
		return 0, errors.NewWithCode(codes.CodeBcrypt, "failed to count cost from hash, %v", err)
	}

	return cost, nil
}
