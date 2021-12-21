package administration

import (
	"fmt"
	"net/http"
)

func (c *Controller) handleAdministration(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	w.WriteHeader(http.StatusOK)
}