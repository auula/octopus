package main

import (
	"github.com/gin-gonic/gin"
	"github.com/higker/octopus/controller"
)

var (
	router *gin.Engine
	port   = ":8080"
)

func init() {
	router = gin.Default()
}

func main() {
	controller.MappingApi(router)
	controller.MappingView(router)
	router.Run(port)
}
