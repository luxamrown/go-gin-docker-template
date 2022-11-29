package api

import (
	"mohamadelabror/gogindocker/usecase"

	"github.com/gin-gonic/gin"
)

type PeopleApi struct {
	peopleListUsecase usecase.GetAllPeopleUsecase
}

func (p *PeopleApi) GetAllPeople() gin.HandlerFunc {
	return func(c *gin.Context) {
		peopleList := p.peopleListUsecase.GetAllPeopleListUsecase()
		c.JSON(200, gin.H{"Data": peopleList})
	}
}

func NewPeopleApi(peopleRoute *gin.RouterGroup, peopleListUsecase usecase.GetAllPeopleUsecase) {
	api := PeopleApi{
		peopleListUsecase: peopleListUsecase,
	}
	peopleRoute.GET("/all", api.GetAllPeople())
}
