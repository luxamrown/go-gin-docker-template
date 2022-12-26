package repository

import "mohamadelabror/gogindocker/model"

type PeopleRepo interface {
	GetAllPeopleList() []model.People
}

type peopleRepo struct {
	people model.People
}

func (p *peopleRepo) GetAllPeopleList() []model.People {
	listPeople := [...]model.People{
		{
			Name: "El",
			Age:  18,
		},
		{
			Name: "Joe",
			Age:  30,
		}, {
			Name: "Julie",
			Age:  29,
		}, {
			Name: "Christian",
			Age:  21,
		},
	}
	return listPeople[:]
}

func NewPeopleRepo(people model.People) PeopleRepo {
	return &peopleRepo{
		people: people,
	}
}
