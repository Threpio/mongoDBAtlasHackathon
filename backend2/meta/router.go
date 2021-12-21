package meta

import "github.com/go-chi/chi"

func (c *Controller) metaRouter() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", c.HandleGetMeta)
		r.Get("/health", c.HandleGetHealth)
	}
}
