package routes

import (
	"github.com/AstroFireWasTaken/ChatAPI/modules/chat"
	"github.com/AstroFireWasTaken/ChatAPI/modules/user"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	user.AddRoutes(router)
	chat.AddRoutes(router)
}
