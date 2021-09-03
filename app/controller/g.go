package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gongyao1992/go-util/helper"
	"net/http"
)

// JSONResult json result
type JSONResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func downloadFilePath(ctx *gin.Context, filePath string)  {
	// 文件下载
	fileName := helper.GetConfigFile(filePath)
	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
	ctx.File(filePath)
}

func Ping(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, JSONResult{
		Code: 0,
		Data: "pong",
	})
	return
}