package auth

import (
	"context"
	"reflect"

	"github.com/golang-jwt/jwt/v5"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

type JWTInterface[C jwt.Claims] interface {
	Generate(ctx context.Context) (string, error)
	Validate(ctx context.Context, tokenString string) (*jwt.Token, error)
	ExtractClaims(ctx context.Context, jwtToken *jwt.Token) (C, error)
	WithSigningMethod(signingMethod jwt.SigningMethod) JWTInterface[C]
}

type jwtimpl[C jwt.Claims] struct {
	secretKey     []byte
	claims        C
	signingMethod jwt.SigningMethod
}

func InitJWT[C jwt.Claims](secretKey []byte, claims C) JWTInterface[C] {
	j := jwtimpl[C]{
		secretKey:     secretKey,
		claims:        claims,
		signingMethod: jwt.SigningMethodHS256,
	}
	return &j
}

func (j *jwtimpl[C]) Generate(ctx context.Context) (string, error) {
	jwtToken := jwt.NewWithClaims(j.signingMethod, j.claims)
	jwtString, err := jwtToken.SignedString(j.secretKey)
	if err != nil {
		return "", errors.NewWithCode(codes.CodeJWTSignedStringError, "cannot signed string, %v", err.Error())
	}

	return jwtString, nil
}

func (j *jwtimpl[C]) Validate(ctx context.Context, tokenString string) (*jwt.Token, error) {
	kind := reflect.TypeOf(j.claims).Kind()

	switch kind {
	case reflect.Pointer:
		keyFunc := func(jwtToken *jwt.Token) (any, error) {
			if _, isOk := jwtToken.Method.(*jwt.SigningMethodHMAC); !isOk {
				return nil, errors.NewWithCode(codes.CodeJWTInvalidMethod, "invalid token method algoritm")
			}
			return j.secretKey, nil
		}

		jwtToken, err := jwt.ParseWithClaims(tokenString, j.claims, keyFunc)
		if err != nil {
			return nil, errors.NewWithCode(codes.CodeJWTParseWithClaimsError, "cannot parse token with claims, %v", err)
		}

		return jwtToken, nil
	default:
		return nil, errors.NewWithCode(codes.CodeJWTInvalidClaimsType, "claims type must be a pointer but got %v", kind.String())
	}
}

func (j *jwtimpl[C]) ExtractClaims(ctx context.Context, jwtToken *jwt.Token) (C, error) {
	claims, isOk := jwtToken.Claims.(C)
	if !isOk {
		return j.claims, errors.NewWithCode(codes.CodeJWTInvalidClaimsType, "claims type is not equals")
	}

	return claims, nil
}

func (j *jwtimpl[C]) WithSigningMethod(signingMethod jwt.SigningMethod) JWTInterface[C] {
	j.signingMethod = signingMethod
	return j
}
