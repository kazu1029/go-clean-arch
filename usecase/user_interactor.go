package usecase

import (
	"github.com/kazu1029/go-clean-arch/domain"
	"github.com/kazu1029/go-clean-arch/usecase/interfaces"
)

type UserInteractor struct {
	UserRepository interfaces.UserRepository
	Logger         interfaces.Logger
}

func (i *UserInteractor) Add(u domain.User) (int, error) {
	i.Logger.Log("store user!")
	return i.UserRepository.Store(u)
}
