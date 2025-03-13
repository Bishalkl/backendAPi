package user

import (
	"net/http"

	"github.com/bishalkl/backendAPi/types"
	"github.com/bishalkl/backendAPi/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// Handler for  login
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}

// Handler for register
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// check if the user exists
	
	// if it doesn't we create the new user

}
