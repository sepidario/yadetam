package web

import (
	"encoding/json"
	"net/http"

	"github.com/sepidario/yadetam/yadetam"
)

// UserHandler is an implemation of user web API
type UserHandler struct {
	us yadetam.UserService
}

// NewUserHandler create a new user handler
func NewUserHandler(us yadetam.UserService) *UserHandler {
	return &UserHandler{us}
}

// Register a new user
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	user := &yadetam.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.us.CreateUser(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
