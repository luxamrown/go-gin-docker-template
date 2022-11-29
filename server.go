package main

import (
	"mohamadelabror/gogindocker/delivery/api"
	"mohamadelabror/gogindocker/manager"

	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine *gin.Engine
}

func (a *appServer) initHandlers() {
	a.v1()
}

func (a *appServer) v1() {
	peopleApiGroup := a.routerEngine.Group("/people")
	UseCaseManager := manager.NewUseCaseManager(manager.NewRepoManager())
	api.NewPeopleApi(peopleApiGroup, UseCaseManager.GetAllPeopleUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.routerEngine.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func Server() AppServer {
	r := gin.Default()
	return &appServer{
		routerEngine: r,
	}
}
