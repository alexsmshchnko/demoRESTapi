package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID        string
	Firstname string
	Lastname  string
	Email     string
	Age       uint
	Created   time.Time
}

func getUser(ctx *gin.Context) {
	u := pg.getUser(ctx.Params.ByName("id"))

	if u.ID == "" {
		ctx.IndentedJSON(http.StatusNotFound, "")
	} else {
		ctx.IndentedJSON(http.StatusOK, u)
	}
}

func addUser(ctx *gin.Context) {
	var err error
	u := &User{}
	if err = ctx.BindJSON(u); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err = pg.addUser(u); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, u)
	}
}

func patchUser(ctx *gin.Context) {
}

type db struct {
	*sql.DB
}

var pg db

func connectDB() (db db, err error) {
	const (
		host     = "localhost"
		port     = 5430
		user     = "postgres_user"
		password = "postgres_password"
		dbname   = "postgres_db"
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	if db.DB, err = sql.Open("postgres", psqlconn); err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	return
}

func (pg *db) getUser(id string) *User {
	row := pg.QueryRow("SELECT id, firstname, lastname, email, age, created FROM public.user WHERE id = $1", id)

	u := &User{}
	if err := row.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Age, &u.Created); err != nil {
		fmt.Println(err)
	}

	return u
}

func (pg *db) addUser(u *User) (err error) {
	if err = pg.QueryRow("INSERT INTO public.user (id, firstname, lastname, email, age, created) VALUES(gen_random_uuid(), $1, $2, $3, $4, now()) RETURNING id, created",
		u.Firstname, u.Lastname, u.Email, u.Age).Scan(&u.ID, &u.Created); err != nil {
		fmt.Println(err)
	}
	return
}

func main() {
	var err error
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	router.GET("/user/:id", getUser)
	router.POST("/users", addUser)
	router.PATCH("/user/:id", patchUser)

	if pg, err = connectDB(); err != nil {
		fmt.Println("connectDB error:", err)
		os.Exit(1)
	}

	router.Run("localhost:8080")
}
