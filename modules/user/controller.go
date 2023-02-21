package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func handleCreateUser(ctx *gin.Context) {
	var user User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = CreateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func handleGetUsers(ctx *gin.Context) {
	users, err := FetchUsers()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func handleGetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := FetchUserById(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if user == nil {
		ctx.AbortWithError(http.StatusNotFound, errors.New("No user found with id "+idStr))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func handleUpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user User
	err = ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user.ID = id

	err = UpdateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func handleDeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = DeleteUser(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}

func AddRoutes(router *gin.Engine) {
	router.POST("/users", handleCreateUser)
	router.GET("/users", handleGetUsers)
	router.GET("/users/:id", handleGetUserById)
	router.PUT("/users/:id", handleUpdateUser)
	router.DELETE("/users/:id", handleDeleteUser)
}
