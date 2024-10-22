package main

import (
	"fmt"
	"gotu-bookstore/cmd/gotu-bookstore/configs"
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/cmd/gotu-bookstore/routes"
	"os"

	"gotu-bookstore/pkg/logger"

	"github.com/gin-gonic/gin"
)

type EngineBuilder struct {
	appConfig configs.AppConfig
	engine    *gin.Engine
}

func NewEngineBuilder(appConfig configs.AppConfig) *EngineBuilder {
	appMode := appConfig.Server.AppMode
	if appMode == constants.AppDevelopmentMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	return &EngineBuilder{
		appConfig: appConfig,
		engine:    gin.New(),
	}
}

func (n *EngineBuilder) SetProxy() *EngineBuilder {
	appMode := n.appConfig.Server.AppMode
	if appMode != constants.AppDevelopmentMode {
		n.engine.SetTrustedProxies(nil)
	}
	return n
}

func (n *EngineBuilder) RegisterMiddleware(handlerFunc gin.HandlerFunc) *EngineBuilder {
	n.engine.Use(handlerFunc)
	return n
}

func (n *EngineBuilder) RegisterNoRoute(handlerFunc gin.HandlerFunc) *EngineBuilder {
	n.engine.NoRoute(handlerFunc)
	return n
}

func (n *EngineBuilder) InitRoutes() *EngineBuilder {
	routes.InitRoutes(n.engine)
	return n
}

func (n *EngineBuilder) Run() {
	err := n.engine.Run(fmt.Sprintf("%s:%d", n.appConfig.Server.AppHost, n.appConfig.Server.AppPort))
	if err != nil {
		log := logger.LogInstance
		log.Fatalf("main-service has been stopped with error %s", err.Error())
		os.Exit(1)
	}
}
