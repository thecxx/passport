package mysql

type UserTable struct {
	connector interface{}
}

func NewUserTable(connector interface{}) *UserTable {
	return &UserTable{connector}
}
