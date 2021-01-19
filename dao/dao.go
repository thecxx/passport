package dao

import (
	"database/sql"
)

var (
	// External resources on which the service depends.
	defaultConnection *sql.Conn = nil

	// Data access object.
	Account *AccountDao = nil
)

// Initialize dao package.
func Init() {
	Account = NewAccountDao(defaultConnection)
}
