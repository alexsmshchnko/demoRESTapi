package main

import (
	"fmt"
	"net/http"
	"os"

	repo "demorestapi/internal/adapters/repository"
	log "demorestapi/internal/common/logs"
	ports "demorestapi/internal/ports"
	service "demorestapi/internal/service"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"moul.io/chizap"
)

func main() {
	pg, err := repo.ConnectDB()
	if err != nil {
		fmt.Println("connectDB error:", err)
		os.Exit(1)
	}

	l := log.NewLogger()

	srv := service.NewService(&pg, &pg, l)

	h := ports.NewHttpServer(srv)
	router := chi.NewRouter()

	router.Use(chizap.New(l.Logger, &chizap.Opts{
		WithReferer:   true,
		WithUserAgent: true,
	}))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) })

	router.Get("/user/{id}", h.GetUser)
	router.Post("/users", h.AddUser)
	router.Patch("/user/{id}", h.PatchUser)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
