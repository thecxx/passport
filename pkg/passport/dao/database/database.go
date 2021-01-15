package database

import (
	"github.com/thecxx/passport/pkg/passport/dao/database/mysql"
)

type DefaultDatabase struct {
	User *mysql.UserTable
}

var (
	Default *DefaultDatabase
)
