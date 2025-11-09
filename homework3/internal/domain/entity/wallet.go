package entity


import (
	"fmt"
	"strconv"
)

func NewWallet() *wallet {
	wall := new(wallet)
	return wall
}

type wallet struct {
	BalanceInRuble int `json:"balanceInRuble`
	BalanceInKopeck int8 `json:"balanceInKopeck"`
}

func (w *wallet) GetKopeck() string {
	return strconv.Itoa(int(w.BalanceInKopeck))
}

func (w *wallet) GetRuble() string {
	return strconv.Itoa(w.BalanceInRuble)
}

func (w *wallet) TopUpBalance(ruble int, kopeck int8) {
	w.BalanceInRuble += ruble
	w.BalanceInRuble += (int(w.BalanceInKopeck) + int(kopeck)) / 100
	w.BalanceInKopeck = int8((int(w.BalanceInKopeck) + int(kopeck)) % 100)
}

func (w *wallet) WithdrawMoney(ruble int, kopeck int8) error {
	newRuble := w.BalanceInRuble - ruble
	newKopeck := w.BalanceInKopeck - kopeck
	if (newKopeck < 0) {
		newRuble--
		newKopeck = 100 + newKopeck
	}
	if (newRuble < 0) {
		return fmt.Errorf("insufficient funds in the account to write off")
	}
	w.SetBalance(newRuble, newKopeck)
	return nil
}

func (w *wallet) SetBalance(ruble int, kopeck int8) {
	w.BalanceInKopeck = kopeck
	w.BalanceInRuble = ruble
}