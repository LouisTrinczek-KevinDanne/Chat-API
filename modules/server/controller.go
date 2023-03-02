package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func handleCreateServer(ctx *gin.Context) {
	var server Server
	err := ctx.BindJSON(&server)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = CreateServer(&server)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, server)
}

func handleGetServers(ctx *gin.Context) {
	servers, err := FetchServers()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, servers)
}

func handleGetServerById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	server, err := FetchServerByID(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if server == nil {
		ctx.AbortWithError(http.StatusNotFound, errors.New("No server found with id "+idStr))
		return
	}

	ctx.JSON(http.StatusOK, server)
}

func handleUpdateServer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var server Server
	err = ctx.BindJSON(&server)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	server.ID = id

	err = UpdateServer(&server)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, server)
}

func handleDeleteServer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = DeleteServer(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}

func AddRoutes(router *gin.Engine) {
	router.POST("/servers", handleCreateServer)
	router.GET("/servers", handleGetServers)
	router.GET("/servers/:id", handleGetServerById)
	router.PUT("/servers/:id", handleUpdateServer)
	router.DELETE("/servers/:id", handleDeleteServer)
}
