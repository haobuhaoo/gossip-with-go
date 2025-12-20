package comments

import "github.com/go-chi/chi/v5"

// Routes group all comment related HTTP endpoints together, with the base prefix path /comments.
// It connects the URLS to their respective handler methods.
func Routes(router chi.Router, h *handler) {
	router.Route("/comments", func(r chi.Router) {
		r.Get("/all/{postId}", h.FindCommentsByPost)
		r.Post("/", h.CreateComment)
		r.Put("/{id}", h.UpdateComment)
		r.Delete("/{id}", h.DeleteComment)
	})
}
