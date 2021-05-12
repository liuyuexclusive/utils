package jwt

import (
	"testing"
)

const (
	testKey = "sadhasldjkko126312jljdkhfasu0"
	testId  = "testUser"
)

func TestJwt(t *testing.T) {
	token, err := GenToken(testId, testKey)
	if err != nil {
		t.Error(err)
	}

	t.Log(token)

	claims, err := GetClaims(token, testKey)

	if err != nil {
		t.Error(err)
	}

	t.Log(claims)

	id := claims["jti"]
	if id != testId {
		t.Errorf("got: %s,want: %s", id, testId)
	}
}
