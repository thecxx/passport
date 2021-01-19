package schema

type Account struct {
	HashCode string `column:"hashcode"`
	Type     string `column:"type"`
	Name     string `column:"name"`
	UserId   string `column:"uid"`
	Status   int    `column:"status"`
}

// AccountUserId returns a uid column.
func AccountUserId(value string) Column {
	return Column{Name: "uid", Value: value}
}

// AccountStatus returns a status column.
func AccountStatus(value int) Column {
	return Column{Name: "status", Value: value}
}
