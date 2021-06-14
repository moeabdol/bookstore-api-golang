package utils

import (
	"errors"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Payload type struct
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"iat"`
	ExpiresAt time.Time `json:"exp"`
}

// Different types of errors returned by the VerifyToken function
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// newPayload function creates a new payload with specific username and duration
func newPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}

// CreateToken function to create a new JWT token for a specific username and
// duration
func CreateToken(username string, duration time.Duration, privateKeyPath string) (string, error) {
	payload, err := newPayload(username, duration)
	if err != nil {
		return "", err
	}

	privateKeyBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	jwtString, err := jwtToken.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return "Bearer " + jwtString, nil
}

// VerifyToken function checks if the token is valid or not
func VerifyToken(token string, publicKeyPath string) (*Payload, error) {
	token = token[7:]

	publicKeyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseECPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, err
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, ErrInvalidToken
		}
		return publicKey, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
