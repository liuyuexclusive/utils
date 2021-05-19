package handler

import (
	"context"
	"errors"

	"github.com/yuexclusive/utils/crypto"
	"github.com/yuexclusive/utils/db"
	"github.com/yuexclusive/utils/jwt"
	"github.com/yuexclusive/utils/srv/auth/proto/auth"
	"github.com/yuexclusive/utils/srv/basic/model"
)

type Handler struct{}

const (
	mySigningKey = "sadhasldjkko126312jljdkhfasu0"
)

func getToken(id, key string) (string, error) {
	var user model.User

	conn, err := db.Open()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	conn.Where("name=?", id).First(&user)

	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", errors.New("invalid user")
	}

	if key == "" {
		return "", errors.New("please input the password")
	}

	pwd := crypto.Sha256(key + user.Salt)

	if pwd != user.Pwd {
		return "", errors.New("wrong password")
	}

	return jwt.GenToken(id, mySigningKey)
}

func (e *Handler) Auth(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {
	token, err := getToken(req.Id, req.Key)
	if err != nil {
		return nil, err
	}
	return &auth.AuthResponse{Token: token}, nil
}

func (e *Handler) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	claims, err := jwt.GetClaims(req.Token, mySigningKey)

	if err != nil {
		return nil, err
	}

	return &auth.ValidateResponse{Name: claims["jti"].(string)}, nil
}
