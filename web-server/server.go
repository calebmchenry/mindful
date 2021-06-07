package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
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

	a.Router = r
}

func (a *App) Serve() {
	http.ListenAndServe(":8080", a.Router)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
