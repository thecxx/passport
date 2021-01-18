package service

import (
	"context"
	"testing"

	"github.com/thecxx/passport/pkg/passport/dao"
)

func TestAccountService_Login(t *testing.T) {

	dao.Init()
	Init()

	err := Account.Login(context.Background(), "username", "test", "password", map[string]string{})
	if err != nil {
		t.Errorf("login failed, err = %s\n", err)
	} else {
		t.Logf("ok")
	}
}
