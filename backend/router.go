package backend

import (
	"github.com/gorilla/mux"
)

// Creates a router that handles the routes specified in the routes.go file
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
