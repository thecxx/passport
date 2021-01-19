package schema

type UserId struct {
	UserId     string `column:"uid"`
	Validators int    `column:"validators"`
	RefCount   int    `column:"refcount"`
}

// UserIdValidators returns a validators column.
func UserIdValidators(value int) Column {
	return Column{Name: "validators", Value: value}
}

// UserIdRefCount returns a refcount column.
func UserIdRefCount(value int) Column {
	return Column{Name: "refcount", Value: value}
}
