package product

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Hanlder type
type Handler struct {
}

// for instance
func NewHandler() *Handler {
	return &Handler{}
}

// Handler to handler to route
func (p *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/products", p.getProducts).Methods("GET")
	router.HandleFunc("/products/{id}", p.getProduct).Methods("GET")
}

// Handler to get all the product
func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get all the products")
}

// Handler to get specific product
func (h *Handler) getProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get specific the product")
}
