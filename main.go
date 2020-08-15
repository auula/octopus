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
	controller.NewView(router)
	controller.NewApi(router)
	router.Run(port)
}
