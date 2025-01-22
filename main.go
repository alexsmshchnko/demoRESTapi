package main

import (
	"fmt"
	"net/http"
	"os"

	repo "demorestapi/internal/adapters/repository"
	ports "demorestapi/internal/ports"
	service "demorestapi/internal/service"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

// func setMiddlewares(router *chi.Mux) {
// 	router.Use(middleware.RequestID)
// 	router.Use(middleware.RealIP)
// 	router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
// 	router.Use(middleware.Recoverer)

// 	addCorsMiddleware(router)
// 	addAuthMiddleware(router)

// 	router.Use(
// 		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
// 		middleware.SetHeader("X-Frame-Options", "deny"),
// 	)
// 	router.Use(middleware.NoCache)
// }

func main() {
	pg, err := repo.ConnectDB()
	if err != nil {
		fmt.Println("connectDB error:", err)
		os.Exit(1)
	}

	srv := service.NewService(&pg, &pg)

	h := ports.NewHttpServer(srv)

	router := chi.NewRouter()
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) })

	// l := log.NewLogger()
	// rl:= middleware.RequestLogger(  )
	// router.Use(rl)

	router.Get("/user/{id}", h.GetUser)
	router.Post("/users", h.AddUser)
	router.Patch("/user/{id}", h.PatchUser)

	http.ListenAndServe("localhost:8080", router)
}
