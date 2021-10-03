package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

//var router *chi.Mux
//var db *sql.DB

func Routers() chi.Router {

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Home Page"))
	})

	r.Mount("/user", userRouter())

	return r
}