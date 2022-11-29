package usecase

import (
	"mohamadelabror/gogindocker/model"
	"mohamadelabror/gogindocker/repository"
)

type GetAllPeopleUsecase interface {
	GetAllPeopleListUsecase() []model.People
}

type getAllPeopleUsecase struct {
	peopleRepo repository.PeopleRepo
}

func (p *getAllPeopleUsecase) GetAllPeopleListUsecase() []model.People {
	return p.peopleRepo.GetAllPeopleList()
}

func NewPeopleUsecase(peopleRepo repository.PeopleRepo) GetAllPeopleUsecase {
	return &getAllPeopleUsecase{
		peopleRepo: peopleRepo,
	}
}
