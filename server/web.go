package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"micro/upperspective/app/route"
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

	// 复购
	route.AddFugou(group)
	// 费用数据
	route.AddFee(group)
	// 客户数据
	route.AddClient(group)
	// 业务员数据
	route.AddSaleman(group)

	router.Run(":" + *port)
}