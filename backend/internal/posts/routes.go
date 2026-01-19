package posts

import "github.com/go-chi/chi/v5"

// Routes group all post related HTTP endpoints together, with the base prefix path /posts.
// It connects the URLS to their respective handler methods.
func Routes(router chi.Router, h *handler) {
	router.Route("/posts", func(r chi.Router) {
		r.Get("/all/{topicId}", h.FindPostsByTopic)
		r.Get("/{topicId}/search", h.SearchPost)
		r.Get("/{topicId}/{postId}", h.FindPostByID)
		r.Post("/", h.CreatePost)
		r.Put("/{id}", h.UpdatePost)
		r.Delete("/{id}", h.DeletePost)
	})
}
