package routes

import (
	"github.com/LouisTrinczek-KevinDanne/Chat-API/modules/chat"
	"github.com/LouisTrinczek-KevinDanne/Chat-API/modules/user"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	user.AddRoutes(router)
	chat.AddRoutes(router)
}
