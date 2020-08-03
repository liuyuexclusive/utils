package jwt

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const mySigningKey string = "sadhasldjkko126312jljdkhfasu0"

func Sha512(s string, salt string) string {
	h := sha512.New()
	h.Write([]byte(s + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func GetToken(id string) (string, error) {
	mySigningKey := []byte(mySigningKey)

	if id == "" {
		return "", errors.New("无效的id")
	}

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    "future",
		Id:        id,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetClaims(token string) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	jt, err := jwt.ParseWithClaims(
		token, &claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !jt.Valid {
		return nil, errors.New("无效token")
	}

	return claims, nil
}
