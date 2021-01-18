package dao

import (
	"context"
	"database/sql"

	"github.com/thecxx/passport/pkg/passport/errors"

	"github.com/thecxx/passport/pkg/passport/dao/database/mysql/schema"

	"github.com/thecxx/passport/pkg/passport/dao/database/mysql"
)

type AccountProtected struct {
	schema.Account
}

type AccountDao struct {
	at *mysql.AccountTable
}

// New account dao.
func NewAccountDao(defaultConn *sql.Conn) *AccountDao {
	return &AccountDao{
		at: mysql.NewAccountTable(defaultConn),
	}
}

func (a *AccountDao) ProbeAccount(ctx context.Context, atype, aname string) (bool, error) {
	return true, nil
}

// CreateAccount creates a new account and returns the hashcode.
func (a *AccountDao) CreateAccount(ctx context.Context, atype, aname string, uid string, auth int, secret, salt string) (string, error) {
	account := schema.Account{
		HashCode: a.hashcode(atype, aname),
		Type:     atype,
		Name:     aname,
		Auth:     auth,
		Secret:   secret,
		Salt:     salt,
		UserId:   uid,
		Status:   0,
	}
	err := a.at.CreateAccount(ctx, account)
	if err != nil {
		return "", err
	}
	return account.HashCode, nil
}

func (a *AccountDao) GetAccount(ctx context.Context, atype, aname string) (AccountProtected, error) {

	exists, err := a.ProbeAccount(ctx, atype, aname)
	if err != nil {
		return AccountProtected{}, err
	}
	if !exists {
		return AccountProtected{}, errors.ErrorAccountNotFound
	}
	account, err := a.at.GetAccount(ctx, a.hashcode(atype, aname))
	if err != nil {
		return AccountProtected{}, err
	}
	return AccountProtected{account}, nil
}

func (a *AccountDao) hashcode(atype, aname string) string {
	return ""
}
