package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - store pointer to our services
type Handler struct {
	Router *mux.Router
}

// Newhandler - returns a pointer to a handler
func Newhandler() *Handler {
	return &Handler{}
}

// SetupRoutes - sets up all the router for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!\n")
	})
}
