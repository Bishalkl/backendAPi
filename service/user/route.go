package user

import (
	"net/http"

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
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	
}

// Handler for register
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}
