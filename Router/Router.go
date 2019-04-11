package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter() *mux.Router {
	// cors 配置
	c := cors.New(cors.Options{
		// AllowedOrigins:   []string{"foo.com"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowCredentials: true,
	})
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		// cors 实现
		handler := c.Handler(route.HandlerFunc)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
