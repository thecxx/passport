package schema

type Account struct {
	HashCode string `column:"hashcode"`
	Type     string `column:"type"`
	Name     string `column:"name"`
	Auth     int    `column:"auth"`
	Salt     string `column:"salt"`
	Secret   string `column:"secret"`
	UserId   string `column:"uid"`
	Status   int    `column:"status"`
}

// AccountHashCode returns a hashcode column.
func AccountHashCode(value string) Column {
	return Column{Name: "hashcode", Value: value}
}

// AccountType returns a type column.
func AccountType(value string) Column {
	return Column{Name: "type", Value: value}
}

// AccountName returns a name column.
func AccountName(value string) Column {
	return Column{Name: "name", Value: value}
}

// AccountAuth returns a auth column.
func AccountAuth(value int) Column {
	return Column{Name: "auth", Value: value}
}

// AccountSalt returns a salt column.
func AccountSalt(value string) Column {
	return Column{Name: "salt", Value: value}
}

// AccountSecret returns a secret column.
func AccountSecret(value string) Column {
	return Column{Name: "secret", Value: value}
}

// AccountUserId returns a uid column.
func AccountUserId(value string) Column {
	return Column{Name: "uid", Value: value}
}

// AccountStatus returns a status column.
func AccountStatus(value int) Column {
	return Column{Name: "status", Value: value}
}
