package ingest

import (
	"encoding/json"
	"fmt"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/types"
	"io/ioutil"
	"net/http"
)

func (c *Controller) handleIngest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) handleStructuredIngest(w http.ResponseWriter, r *http.Request) {

	var structuredIngestRequest types.StructuredIngestRequest
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &structuredIngestRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.structuredIngest(&structuredIngestRequest); err != nil {
		c.logger.Debug(fmt.Sprintf("Error in structured ingest: %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) handleSearchIngest(w http.ResponseWriter, r *http.Request) {

	var searchIngestRequest types.SearchIngestRequest
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &searchIngestRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//if err := c.searchIngest(&searchIngestRequest); err != nil {
	//	c.logger.Debug(fmt.Sprintf("Error in search ingest: %s", err))
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

