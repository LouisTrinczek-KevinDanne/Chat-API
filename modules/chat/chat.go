package chat

import (
	"github.com/AstroFireWasTaken/ChatAPI/modules/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Channels     []Channel   `json:"channels"`
	Members      []user.User `json:"members"`
	CreationTime time.Time   `json:"creationTime"`
}

type Channel struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Messages []ChannelMessage `json:"messages"`
}

type ChannelMessage struct {
	ID        int       `json:"id"`
	Sender    user.User `json:"sender"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func handleGetServers(ctx *gin.Context) {
	servers := make([]Server, 0)
	ctx.JSON(http.StatusOK, servers)
}

func AddRoutes(router *gin.Engine) {
	router.GET("/servers", handleGetServers)
}
