package main

import (
	"github.com/AstroFireWasTaken/ChatAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Init(router)
	router.Run()
}
