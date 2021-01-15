package service

var (
	User    *UserService
	Account *AccountService
)

// Initialize service package.
func Init() {
	User = NewUserService()
	Account = NewAccountService()
}
