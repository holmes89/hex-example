package user

type UserRepo interface {
	CreateAccount(account *Account) error
	GetUser(username string) (*Account, error)
}