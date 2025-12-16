package topics

import "github.com/go-chi/chi/v5"

// Routes group all topic related HTTP endpoints together, with the base prefix path /topics.
// It connects the URLS to their respective handler methods.
func Routes(router chi.Router, h *handler) {
	router.Route("/topics", func(r chi.Router) {
		r.Get("/", h.ListTopics)
		r.Get("/{id}", h.FindTopicByID)
		r.Post("/", h.CreateTopic)
		r.Put("/{id}", h.UpdateTopic)
		r.Delete("/{id}", h.DeleteTopic)
	})
}
