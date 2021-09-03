package route

import (
	"github.com/gin-gonic/gin"
	"micro/upperspective/app/controller"
	"micro/upperspective/app/middleWare"
)

func AddTest(g *gin.RouterGroup)  {
	rg := g.Group("/api/test")

	rg.Use(middleWare.MiddleWare())

	rg.GET("/ping", controller.Ping)
}