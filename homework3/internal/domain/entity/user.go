package entity

import (
	"fmt"
	"strconv"
)

type User struct {
	Name string `json:"name"`
	Wall *wallet `json:"wallet"`
	PhoneNumber string `json:"phoneNumber"`
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetPhoneNumber() string {
	return u.PhoneNumber
}

func (u *User) GetRuble() string {
	return u.Wall.GetRuble()
}

func (u *User) GetKopeck() string {
	return u.Wall.GetKopeck()
}

func validatePhoneNumber(phoneNumber string) error {
	runeNumber := []rune(phoneNumber)
	if (len(runeNumber) != 16) {
		return fmt.Errorf("invalid number format")
	}
	for i := 0; i < 16; i++ {
		switch i {
		case 0:
			if (string(runeNumber[i]) != "+") {
				return fmt.Errorf("invalid number format")
			}
		case 2:
			if (string(runeNumber[i]) != "-") {
				return fmt.Errorf("invalid number format")
			}
		case 6:
			if (string(runeNumber[i]) != "-") {
				return fmt.Errorf("invalid number format")
			}
		case 10:
			if (string(runeNumber[i]) != "-") {
				return fmt.Errorf("invalid number format")
			}
		case 13:
			if (string(runeNumber[i]) != "-") {
				return fmt.Errorf("invalid number format")
			}
		default:
			if _, err := strconv.Atoi(string(runeNumber[i])); err != nil {
				return fmt.Errorf("invalid number format")
			}				
		}
	}
	return nil
}

func (u *User) setNumber(number string) error {
	err := validatePhoneNumber(number)
	if (err != nil) {
		return err
	}
	u.PhoneNumber = number
	return nil
}

func validateName(name string) error {
	if (name == "") {
		return fmt.Errorf("invalid name format")
	}
	return nil
}

func (u *User) setName(name string) error {
	err := validateName(name)
	if (err != nil) {
		return err
	}
	u.Name = name
	return nil
}

func NewUser(name string, number string) (*User, error){
	person := new(User)
	err := person.setName(name)
	if (err != nil) {
		return nil, err
	}
	err = person.setNumber(number)
	if (err != nil) {
		return nil, err
	}
	person.Wall = NewWallet()
	return person, nil
}

func (u *User) setBalance(balanceInRuble int, balanceInKopeck int8) error {
	if (balanceInRuble < 0 || balanceInKopeck < 0) {
		return fmt.Errorf("invalid replenishment format")
	}
	if (balanceInKopeck >= 100) {
		return fmt.Errorf("invalid replenishment format")
	}
	u.Wall.SetBalance(balanceInRuble, balanceInKopeck)
	return nil
}

func GetUser(name string, balanceInRuble int, balanceInKopeck int8, number string) (*User, error) {
	person := new(User)
	err := person.setName(name)
	if (err != nil) {
		return nil, err
	}
	err = person.setNumber(number)
	if (err != nil) {
		return nil, err
	}
	err = person.setBalance(balanceInRuble, balanceInKopeck)
	if (err != nil) {
		return nil, err
	}
	return person, nil
}

func validateReplenishment(ruble int, kopeck int8) error {
	if (ruble < 0 || kopeck < 0) {
		return fmt.Errorf("invalid replenishment format")
	}
	if (ruble + int(kopeck) == 0) {
		return fmt.Errorf("invalid replenishment format")
	}
	if (kopeck >= 100) {
		return fmt.Errorf("invalid replenishment format")
	}
	return nil
}

func (u *User) TopUpBalance(ruble int, kopeck int8) error {
	err := validateReplenishment(ruble, kopeck)
	if (err != nil) {
		return err
	}
	u.Wall.TopUpBalance(ruble, kopeck)
	return nil
}

func (u *User) WithdrawMoneyFromTheBalance(ruble int, kopeck int8) error {
	err := validateReplenishment(ruble, kopeck)
	if (err != nil) {
		return err
	}
	err = u.Wall.WithdrawMoney(ruble, kopeck)
	if (err != nil) {
		return err
	}
	return nil
}