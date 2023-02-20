package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Password     string    `json:"password,omitempty"`
	CreationTime time.Time `json:"creationTime"`
}

func handleCreateUser(ctx *gin.Context) {
	var user User
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Println("Error while parsing input", err)
		return
	}

	err = CreateUser(&user)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while creating user", err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func handleGetUsers(ctx *gin.Context) {
	users, err := FetchUsers()
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while fetching users", err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func handleGetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		log.Println("Error while parsing id", err)
		return
	}

	user, err := FetchUserById(id)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while fetching user", err)
		return
	}

	if user == nil {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func handleUpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		log.Println("Error while parsing id", err)
		return
	}

	user, err := FetchUserById(id)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		log.Println("Error while fetching user with id", id)
		return
	}

	if user == nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		log.Println("No user found with id", id)
		return
	}

	err = ctx.BindJSON(user)
	if err != nil {
		log.Println("Error while parsing input", err)
		return
	}

	err = UpdateUser(user)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while updating user", err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func handleDeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		log.Println("Error while parsing id", err)
		return
	}

	user, err := FetchUserById(id)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while fetching user with id", id)
		return
	}

	if user == nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		log.Println("No user found with id", id)
		return
	}

	err = DeleteUser(id)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while deleting user", err)
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
