package main

import (
	"fmt"
	"net/http"
	"os"

	repo "demorestapi/internal/adapter/repository"
	"demorestapi/internal/entity"
	service "demorestapi/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func getUser(ctx *gin.Context) {
	u, err := srv.GetUser(ctx.Params.ByName("id"))

	if u == nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
	} else if u.ID == "" {
		ctx.IndentedJSON(http.StatusNotFound, "")
	} else {
		ctx.IndentedJSON(http.StatusOK, u)
	}
}

func addUser(ctx *gin.Context) {
	var err error
	u := entity.NewUser()
	if err = ctx.BindJSON(u); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err = srv.AddUser(u); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, u)
	}
}

func patchUser(ctx *gin.Context) {
	var err error
	u := entity.NewUser()
	if err = ctx.BindJSON(u); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	u.ID = ctx.Params.ByName("id")

	if err = srv.UpdateUser(u); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, u)
	}
}

var (
	pg  repo.Repository
	srv *service.Service
)

func main() {
	var err error
	if pg, err = repo.ConnectDB(); err != nil {
		fmt.Println("connectDB error:", err)
		os.Exit(1)
	}

	srv = service.NewService(&pg, &pg)

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	router.GET("/user/:id", getUser)
	router.POST("/users", addUser)
	router.PATCH("/user/:id", patchUser)

	router.Run("localhost:8080")
}
