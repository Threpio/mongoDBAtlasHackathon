package ingest

import (
	"github.com/go-chi/chi"
)

func (c *Controller) ingestRouter() func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/", func(r chi.Router) {
			r.Get("/", c.handleIngest)
			r.Get("/search", c.handleSearchIngest)
			r.Post("/structured", c.handleStructuredIngest)
		})
	}
}
