package main

import (
	"github.com/gin-gonic/gin"

	"goTest/ctrl"
)

func main() {

	mainRoute := gin.Default()

	mainRoute.POST("/test", ctrl.ApiIndex)

	mainRoute.Run()
}
