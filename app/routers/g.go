package routers

import (
	"github.com/gin-gonic/gin"
	"xiangxin/go-base/app/middleWare"
)

func AddTest(g *gin.RouterGroup)  {
	rg := g.Group("/api/test")

	rg.Use(middleWare.MiddleWare())


	rg.GET("/ping", )
}