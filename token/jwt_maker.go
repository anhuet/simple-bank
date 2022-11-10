package token

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			fmt.Println("keyFnc", ok)
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, _ := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	//fmt.Println(err)
	fmt.Println(reflect.TypeOf(jwtToken))
	//if err != nil {
	//	fmt.Println("parse", err)
	//
	//	verr, ok := err.(*jwt.ValidationError)
	//	if ok && errors.Is(verr.Inner, ErrExpiredToken) {
	//		return nil, ErrExpiredToken
	//	}
	//	return nil, ErrInvalidToken
	//}

	payload, ok := jwtToken.Claims.(*Payload)
	fmt.Println("payload ne", payload)
	fmt.Println("Claims", ok)

	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
