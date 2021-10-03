package routes

import (
	"github.com/go-chi/chi"
	"go_rest_test/handlers"
)

func userRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handlers.UserList)
	r.Post("/", handlers.UserCreate)
	r.Put("/{id}", handlers.UserUpdate)
	r.Delete("/{id}", handlers.UserDelete)
	r.Get("/{id}", handlers.UserDetails)

	//r.Route("/{id}", func(r chi.Router) {
	//	r.Use(PostCtx)
	//	r.Get("/", )
	//	r.Put("/", )
	//	r.Delete("/", ete)
	//})

	return r
}

