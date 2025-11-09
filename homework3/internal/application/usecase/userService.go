package usecase

import (
	"homework3/internal/domain/entity"
	"homework3/internal/domain/repository"
)

func NewUserService(repository repository.UserRepositoryInterface) *UserService {
	service := new(UserService)
	service.repository = repository
	return service
}

type UserService struct {
	repository repository.UserRepositoryInterface
}

func (s *UserService) TopUpBalance(phoneNumber string, ruble int, kopeck int8) error {
	user, err := s.repository.GetUser(phoneNumber)
	if (err != nil) {
		return err
	}
	err = user.TopUpBalance(ruble, kopeck)
	if (err != nil) {
		return err
	}
	err = s.repository.SaveUser(user)
	if (err != nil) {
		return err
	}
	return nil
}

func (s *UserService) TransferMoney(fromNumber string, toNumber string, ruble int, kopeck int8) error {
	fromUser, err := s.repository.GetUser(fromNumber)
	if (err != nil) {
		return err
	}
	err = fromUser.WithdrawMoneyFromTheBalance(ruble, kopeck)
	if (err != nil) {
		return err
	}
	err = s.repository.SaveUser(fromUser)
	if (err != nil) {
		return err
	}

	toUser, err := s.repository.GetUser(toNumber)
	if (err != nil) {
		return err
	}
	err = toUser.TopUpBalance(ruble, kopeck)
	if (err != nil) {
		return err
	}
	err = s.repository.SaveUser(toUser)
	if (err != nil) {
		return err
	}

	return nil
}

func (s *UserService) GetBalance(phoneNumber string) (string, string, error) {
	user, err := s.repository.GetUser(phoneNumber)
	if (err != nil) {
		return "", "", err
	}
	ruble := user.GetRuble()
	kopeck := user.GetKopeck()
	return ruble, kopeck, nil
}

func (s *UserService) NewUser(phoneNumber string, name string) error {
	user, err := entity.NewUser(name, phoneNumber)
	if (err != nil) {
		return err
	}
	err = s.repository.SaveUser(user)
	if (err != nil) {
		return err
	}
	return nil
}