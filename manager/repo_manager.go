package manager

import (
	"mohamadelabror/gogindocker/model"
	"mohamadelabror/gogindocker/repository"
)

type RepoManager interface {
	PeopleRepo() repository.PeopleRepo
}

type repoManager struct {
	peopleModel model.People
}

func (r *repoManager) PeopleRepo() repository.PeopleRepo {
	return repository.NewPeopleRepo(model.People{})
}

func NewRepoManager() RepoManager {
	return &repoManager{
		peopleModel: model.People{},
	}
}
