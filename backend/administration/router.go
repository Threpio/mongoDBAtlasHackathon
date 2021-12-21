package administration

import "github.com/go-chi/chi"

func (c *Controller) administrationRouter() func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/", func(r chi.Router) {
			r.Get("/", c.handleAdministration)
			r.Get("/collections", c.handleCollections)
			r.Post("/collections/create", c.handleCreateCollection)
		})
	}
}