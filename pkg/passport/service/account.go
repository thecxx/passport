package service

import (
	"context"
	"errors"
	"sync"

	"github.com/thecxx/passport/pkg/passport/dao"

	"github.com/thecxx/passport/proto"
)

type AccountPublic struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	UserId string `json:"uid"`
}

type Validator interface {
	Validate(atype, aname string, auth int, salt, secret string, vparams map[string]string) (bool, error)
}

type AccountService struct {
	once       sync.Once
	user       proto.User
	vamux      sync.RWMutex
	validators map[string]Validator
}

// New account service.
func NewAccountService() *AccountService {
	return &AccountService{user: nil}
}

func (a *AccountService) SetUserService(u proto.User) {
	a.once.Do(func() { a.user = u })
}

func (a *AccountService) Login(ctx context.Context, atype, aname, vname string, vparams map[string]string) (AccountPublic, error) {

	a.vamux.RLock()
	validator, ok := a.validators[vname]
	a.vamux.RUnlock()
	if !ok {
		return AccountPublic{}, errors.New("validator not found")
	}

	account, err := dao.Account.GetAccount(ctx, atype, aname)
	if err != nil {
		return AccountPublic{}, err
	}
	if account.Name == "" {
		return AccountPublic{}, errors.New("account not found")
	}

	ok, err = validator.Validate(account.Type, account.Name, account.Auth, account.Salt, account.Secret, vparams)
	if err != nil {
		return AccountPublic{}, err
	}
	// Login failed
	if !ok {

	}
	return AccountPublic{
		Type:   account.Type,
		Name:   account.Name,
		UserId: account.UserId,
	}, nil
}

func (a *AccountService) Register(ctx context.Context, atype, aname, uid string) (string, error) {

	var err error

	exists, err := dao.Account.ProbeAccount(ctx, atype, aname)
	if err != nil {
		return "", err
	}
	if exists {
		return "", errors.New("account already exists")
	}

	if uid == "" {
		uid, err = UserId.AllocUserId(ctx)
		if err != nil {
			return "", err
		}
	} else {
		exists, err := UserId.ProbeUserId(ctx, uid)
		if err != nil {
			return "", err
		}
		if !exists {
			return "", errors.New("uid not found")
		}
	}

	//err = dao.Account.CreateAccount(ctx, atype, aname, uid)
	//if err != nil {
	//	return "", err
	//}

	return uid, nil
}

func (a *AccountService) BindValidator(ctx context.Context, vname string, validator Validator) {
	a.vamux.Lock()
	a.validators[vname] = validator
	a.vamux.Unlock()
}
