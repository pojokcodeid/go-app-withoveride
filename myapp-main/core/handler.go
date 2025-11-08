package core

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	Service UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		Service: GetUserService(),
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name string `json:"name"`
	}
	json.NewDecoder(r.Body).Decode(&payload)

	h.Service.CreateUser(payload.Name)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.Service.GetUsers()
	resp, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
