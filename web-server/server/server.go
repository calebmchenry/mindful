package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Server struct {
	Router *mux.Router
}

func New() Server {
	s := Server{}
	r := mux.NewRouter()

	// Middleware
	r.Use(CorsOriginMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(r)

	r.Path("/api/v1/login").Handler(http.HandlerFunc(foo)).Methods(http.MethodPost, http.MethodOptions)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(AuthMiddleware)
	api.Path("/foo").Handler(http.HandlerFunc(foo)).Methods(http.MethodGet, http.MethodOptions)

	s.Router = r
	return s
}

func (a *Server) Serve() {
	http.ListenAndServe(":8080", a.Router)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
