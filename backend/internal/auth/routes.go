package auth

import "github.com/go-chi/chi/v5"

// Routes group all authentication related HTTP endpoints together, with the base prefix
// path /auth.
// It connects the URLS to their respective handler methods.
func Routes(router chi.Router, h *handler) {
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
	})
}
