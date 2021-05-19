package user_handler

import (
	"context"
	"errors"
	"strings"

	"github.com/yuexclusive/utils/crypto"
	"github.com/yuexclusive/utils/db"
	"github.com/yuexclusive/utils/jwt"
	"github.com/yuexclusive/utils/srv/basic/model"
	"github.com/yuexclusive/utils/srv/basic/proto/user"
)

type Handler struct {
}

const (
	mySigningKey = "sadhasldjkko126312jljdkhfasu0"
)

func auth(id, key string) (string, error) {
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

func (e *Handler) Get(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error) {
	var rsp user.GetResponse
	conn, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var user model.User
	conn.Where("name=?", req.Name).First(&user)
	if user.ID == 0 {
		return nil, errors.New("找不到用户 " + req.Name)
	}
	rsp.Name = user.Name
	rsp.Access = strings.Split(user.Access, ",")
	rsp.Avatar = user.Avatar

	return &rsp, nil
}
