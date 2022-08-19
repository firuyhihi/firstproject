package delivery

import (
	"log"

	"ticket.narindo.com/config"
	"ticket.narindo.com/delivery/controller"
	"ticket.narindo.com/manager"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	usecManager manager.UsecaseManager
	engine      *gin.Engine
	host        string
}

func Server() *appServer {
	r := initRouterConfiguration()
	cfg := config.InitConfig()
	infra := manager.InitInfra(&cfg)
	repoManager := manager.InitRepositoryManager(infra)
	usecaseManager := manager.InitUsecasesManager(repoManager)

	host := cfg.ApiConfig.Url
	return &appServer{
		usecManager: usecaseManager,
		engine:      r,
		host:        host,
	}
}

func (a *appServer) Run() {
	a.initControllers()
	log.Println("Running server on ", a.host)
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

func (a *appServer) initControllers() {
	controller.NewTicketController(a.engine, a.usecManager.TicketUseCase())
	controller.InitUserCrudController(a.engine, a.usecManager)
}

func initRouterConfiguration() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(corsConfiguration())
	return router
}

func corsConfiguration() gin.HandlerFunc {
	return cors.Default()
}
