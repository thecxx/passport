package service

var (
	UserId  *UserIdService
	Account *AccountService
)

// Initialize service package.
func Init() {
	UserId = NewUserIdService()
	Account = NewAccountService()
}
