package server

import (
	"fmt"
	"net/http"

	"github.com/calebmchenry/mindful/web-server/auth"
	"github.com/calebmchenry/mindful/web-server/env"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Router *mux.Router
	DB     *mongo.Client
}

func New() Server {
	s := Server{}
	r := mux.NewRouter()

	// Middleware
	r.Use(corsOriginMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(r)

	// Login should not be behind authentication
	r.Path("/api/v1/login").Handler(http.HandlerFunc(foo)).Methods(http.MethodPost, http.MethodOptions)

	// Authentication required
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(auth.Middleware)
	installRoutes(api)

	s.Router = r

	return s
}

func (s *Server) Serve() {
	port := env.GetHttpPort()
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Listing and Serving on localhost:%s\n", port)
	http.ListenAndServe(addr, s.Router)
}

func installRoutes(r *mux.Router) {
	// Auth
	r.Path("/logout").Handler(http.HandlerFunc(bar)).Methods(http.MethodGet, http.MethodOptions)

	// Topics
	r.Path("/topics").Handler(http.HandlerFunc(bar)).Methods(http.MethodPost, http.MethodOptions)
	r.Path("/topics").Handler(http.HandlerFunc(bar)).Methods(http.MethodGet, http.MethodOptions)
	r.Path("/topics/:id").Handler(http.HandlerFunc(bar)).Methods(http.MethodGet, http.MethodOptions)
	r.Path("/topics/:id").Handler(http.HandlerFunc(bar)).Methods(http.MethodPut, http.MethodOptions)
	r.Path("/topics/:id").Handler(http.HandlerFunc(bar)).Methods(http.MethodDelete, http.MethodOptions)
}

func bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func corsOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := env.GetHttpOrigin()
		w.Header().Set("Access-Control-Allow-Origin", origin)
		next.ServeHTTP(w, r)
	})
}
