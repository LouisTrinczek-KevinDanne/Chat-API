package main

import (
	"github.com/LouisTrinczek-KevinDanne/Chat-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Init(router)
	router.Run()
}
