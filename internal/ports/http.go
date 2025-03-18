package ports

import (
	"demorestapi/internal/entity"
	"demorestapi/internal/service"
	"encoding/json"
	"net/http"
	"strings"
)

type HttpServer struct {
	app *service.Service
}

func NewHttpServer(app *service.Service) *HttpServer {
	return &HttpServer{app: app}
}

type Error struct {
	// Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

func JSONError(httpcode int, msg string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpcode)
	json.NewEncoder(w).Encode(
		Error{
			// Code:    &code,
			Message: &msg,
		},
	)
}

// GetUser godoc
//
//	@Summary		Get user
//	@Description	get user by ID
//	@Tags			user
//	@Produce		json
//	@Param			id	path	int	true	"User ID"
//	@Success		200	{object}	entity.User
//	@Failure		400 {object}	Error
//	@Router			/user/{id} [get]
func (h HttpServer) GetUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	res, err := h.app.GetUser(id)
	if err != nil {
		JSONError(http.StatusBadRequest, err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// AddUser godoc
//
//	@Summary		Add new user
//	@Description	add new user record
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body	entity.User	true	"New user"
//	@Success		200	{object}	entity.User
//	@Failure		400 {object}	Error
//	@Router			/users [post]
func (h HttpServer) AddUser(w http.ResponseWriter, r *http.Request) {
	user := entity.NewUser()
	json.NewDecoder(r.Body).Decode(user)

	if err := h.app.AddUser(user); err != nil {
		JSONError(http.StatusBadRequest, err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// AddUser godoc
//
//	@Summary		Update user
//	@Description	update user record
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"User ID"
//	@Param			user	body	entity.User	true	"User data"
//	@Success		200	{object}	entity.User
//	@Failure		400 {object}	Error
//	@Router			/user/{id} [patch]
func (h HttpServer) PatchUser(w http.ResponseWriter, r *http.Request) {
	user := entity.NewUser()
	json.NewDecoder(r.Body).Decode(user)
	user.ID = strings.TrimPrefix(r.URL.Path, "/user/")

	if err := h.app.UpdateUser(user); err != nil {
		JSONError(http.StatusBadRequest, err.Error(), w)
		return
	}
	json.NewEncoder(w).Encode(user)
}
