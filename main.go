package main

import (
	"fmt"
	"net/http"
	"os"

	"demorestapi/internal/entity"
	repo "demorestapi/internal/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func getUser(ctx *gin.Context) {
	u := pg.GetUser(ctx.Params.ByName("id"))

	if u.ID == "" {
		ctx.IndentedJSON(http.StatusNotFound, "")
	} else {
		ctx.IndentedJSON(http.StatusOK, u)
	}
}

func addUser(ctx *gin.Context) {
	var err error
	u := &entity.User{}
	if err = ctx.BindJSON(u); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err = pg.AddUser(u); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, u)
	}
}

func patchUser(ctx *gin.Context) {
	var err error
	u := &entity.User{}
	if err = ctx.BindJSON(u); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	u.ID = ctx.Params.ByName("id")

	if err = pg.UpdateUser(u); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, u)
	}
}

var pg repo.Repository

func main() {
	var err error
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	router.GET("/user/:id", getUser)
	router.POST("/users", addUser)
	router.PATCH("/user/:id", patchUser)

	if pg, err = repo.ConnectDB(); err != nil {
		fmt.Println("connectDB error:", err)
		os.Exit(1)
	}

	router.Run("localhost:8080")
}
