package users

import "github.com/go-chi/chi/v5"

// Routes group all user related HTTP endpoints together, with the base prefix path /users.
// It connects the URLS to their respective handler methods.
func Routes(router chi.Router, h *handler) {
	router.Route("/users", func(r chi.Router) {
		r.Get("/{name}", h.FindUserByName)
		r.Post("/", h.CreateUser)
	})
}
