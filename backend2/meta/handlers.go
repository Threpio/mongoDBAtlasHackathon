package meta

import "net/http"

func (c *Controller) HandleGetMeta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"meta": "meta"}`))
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) HandleGetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"health": "healthy :)"}`))
	w.WriteHeader(http.StatusOK)
}
