package passport

import (
	"context"

	"github.com/thecxx/passport/dao"
)

type Passport struct {
	account *dao.AccountDao

	uniq UniqueId
}

// New returns a passport.
func New() *Passport {
	return &Passport{}
}

func (p *Passport) ProbeAccount(ctx context.Context, name AccountName) (bool, error) {
	return false, nil
}

func (p *Passport) CreateAccount(ctx context.Context, name AccountName, uid string) (Account, error) {

	var err error

	if uid != "" {

	} else {
		uid, err = p.uniq.Generate()
		if err != nil {
			return Account{}, err
		}
	}

	return Account{}, nil
}

func (p *Passport) RemoveAccount(ctx context.Context, name AccountName) (bool, error) {
	return false, nil
}

func (p *Passport) AccessAccount(ctx context.Context, name AccountName, validator string, vparams map[string]string) (Account, error) {
	return Account{}, nil
}
