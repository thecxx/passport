package mysql

import (
	"context"
	"database/sql"

	"github.com/thecxx/passport/pkg/passport/dao/database/mysql/schema"
)

type AccountTable struct {
	connector *sql.Conn
}

func NewAccountTable(connector *sql.Conn) *AccountTable {
	return &AccountTable{connector}
}

func (a *AccountTable) CreateAccount(ctx context.Context, account schema.Account) error {
	return nil
}

func (a *AccountTable) GetAccount(ctx context.Context, hashcode string) (schema.Account, error) {
	return schema.Account{
		HashCode: "85136c79cbf9fe36bb9d05d0639c70c265c18d37",
		Type:     "username",
		Name:     "kami",
		Auth:     0,
		Salt:     "",
		Secret:   "",
		UserId:   "7c4a8d09ca3762af61e5",
	}, nil
}

func (a *AccountTable) UpdateAccount(ctx context.Context, hashcode string, values []schema.Column) error {
	return nil
}

func (a *AccountTable) DeleteAccount(ctx context.Context, hashcode string) error {
	return nil
}

func (a *AccountTable) table(hashcode string) string {
	return ""
}
