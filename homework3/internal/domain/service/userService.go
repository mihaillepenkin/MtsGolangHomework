package service

type UserServiceInterface interface {
	TopUpBalance(phoneNumber string, ruble int, kopeck int8) error
	TransferMoney(fromNumber string, toNumber string, ruble int, kopeck int8) error
	GetBalance(phoneNumber string) (string, string, error)
	NewUser(phoneNumber string, name string) error
}

