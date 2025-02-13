package main

import (
	"net/http"

	pg "demorestapi/internal/adapters/postgres"
	repository "demorestapi/internal/adapters/repository"
	log "demorestapi/internal/common/logs"
	ports "demorestapi/internal/ports"
	service "demorestapi/internal/service"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"moul.io/chizap"
)

func main() {
	l := log.NewLogger()

	db, err := pg.ConnectDB()
	if err != nil {
		l.Logger.Panic("connectDB error", zap.Error(err))
	}

	repo := repository.NewRepo(&db)

	srv := service.NewService(repo, repo, l)

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
