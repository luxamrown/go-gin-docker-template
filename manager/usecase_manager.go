package manager

import (
	"mohamadelabror/gogindocker/usecase"
)

type UsecaseManager interface {
	GetAllPeopleUseCase() usecase.GetAllPeopleUsecase
}

type usecaseManager struct {
	repo RepoManager
}

func (u *usecaseManager) GetAllPeopleUseCase() usecase.GetAllPeopleUsecase {
	return usecase.NewPeopleUsecase(u.repo.PeopleRepo())
}

func NewUseCaseManager(repo RepoManager) UsecaseManager {
	return &usecaseManager{
		repo: repo,
	}
}
