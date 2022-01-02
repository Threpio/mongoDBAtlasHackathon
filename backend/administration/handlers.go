package administration

import (
	"fmt"
	"net/http"
)

func (c *Controller) handleAdministration(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) createCollection(w http.ResponseWriter, r *http.Request) {




	w.Write([]byte("created collection"))
	w.WriteHeader(http.StatusCreated)
}