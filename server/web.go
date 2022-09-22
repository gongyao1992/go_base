package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"micro/upperspective/app/routers"
	"runtime"
)

var (
	port = flag.String("port", "9092", "Port: http port")
)

func main() {
	// 传入的参数
	flag.Parse()
	//fmt.Println(""db.GetConfig())
	runtime.GOMAXPROCS(runtime.NumCPU())
	router := gin.Default()

	group := router.Group("/go")

	routers.AddTest(group)

	router.Run(":" + *port)
}