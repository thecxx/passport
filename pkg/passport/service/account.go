package service

import (
	"context"
	"sync"

	"github.com/thecxx/passport/proto"
)

type AccountService struct {
	once sync.Once
	user proto.User
}

// New account service.
func NewAccountService() *AccountService {
	return &AccountService{user: nil}
}

func (a *AccountService) SetUserService(u proto.User) {
	a.once.Do(func() { a.user = u })
}

func (a *AccountService) Register(ctx context.Context) error {
	return a.user.Create()
}
