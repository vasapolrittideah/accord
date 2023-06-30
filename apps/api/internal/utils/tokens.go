package utils

import (
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/matthewhartstonge/argon2"
	"time"
)

func GenerateToken(ttl time.Duration, privateKey string, userId uuid.UUID) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", errors.New("unable to decode key: " + err.Error())
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decoded)
	if err != nil {
		return "", errors.New("unable to parse key: " + err.Error())
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"sub": userId,
		"iat": now.Unix(),
		"exp": now.Add(ttl).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", errors.New("unable to sign jwt token: " + err.Error())
	}

	return token, nil
}

func ValidateToken(token string, publicKey string) (*jwt.Token, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, errors.New("unable to decode key: " + err.Error())
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, errors.New("unable to parse key: " + err.Error())
	}

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method: " + t.Header["alg"].(string))
		}
		return key, nil
	})
}

func HashRefreshToken(refreshToken string) (string, error) {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(refreshToken))
	if err != nil {
		return "", errors.New("unable to hash refresh token: " + err.Error())
	}

	return string(encoded), nil
}

func VerifyRefreshToken(encoded string, refreshToken string) (bool, error) {
	return argon2.VerifyEncoded([]byte(refreshToken), []byte(encoded))
}
