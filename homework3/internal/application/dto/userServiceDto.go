package dto

type TopUpBalanceInputDto struct {
	PhoneNumber string `json:"phoneNumber"`
	Ruble int `json:"ruble"`
	Kopeck int8 `json:"kopeck"`
}

type TopUpBalanceOutputDto struct {
	Status string `json:"status"`
	Msg string `json:"msg"`
}

type TransferMoneyInputDto struct {
	ToNumber string `json:"toNumber"`
	FromNumber string `json:"fromNumber"`
	Ruble int `json:"ruble"`
	Kopeck int8 `json:"kopeck"`
}

type TransferMoneyOutputDto struct {
	Status string `json:"status"`
	Msg string `json:"msg"`
}

type GetBalanceInputDto struct {
	PhoneNumber string `json:"phoneNumber"`
}

type GetBalanceOutputDto struct {
	Status string `json:"status"`
	Msg string `json:"msg"`
	Ruble string `json:"ruble"`
	Kopeck string `json:"kopeck"`
}

type NewUserInputDto struct {
	PhoneNumber string `json:"phoneNumber"`
	Name string `json:"name"`
}

type NewUserOutputDto struct {
	Status string `json:"status"`
	Msg string `json:"msg"`
}