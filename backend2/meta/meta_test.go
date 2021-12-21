package meta

import (
	"github.com/go-chi/chi"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/db"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_HandleGetHealth(t *testing.T) {

	dbTest, err := db.NewDB()
	if err != nil {
		t.Error(err)
	}
	log := logger.New("DEBUG")

	testController, err := NewController(*dbTest, *log)
	if err != nil {
		t.Error(err)
	}

	rt := chi.NewRouter()
	rt.Get("/", testController.HandleGetHealth)

	testServer := httptest.NewServer(rt)

	// Actual Test
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
	if err != nil {
		t.Error(err)
	}
	_, err = client.Do(req)
	if err != nil {
		t.Error(err)
	}
}

func TestController_HandleGetMeta(t *testing.T) {

	dbTest, err := db.NewDB()
	if err != nil {
		t.Error(err)
	}
	log := logger.New("DEBUG")

	testController, err := NewController(*dbTest, *log)
	if err != nil {
		t.Error(err)
	}

	rt := chi.NewRouter()
	rt.Get("/", testController.HandleGetMeta)

	testServer := httptest.NewServer(rt)

	// Actual Test
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
	if err != nil {
		t.Error(err)
	}
	_, err = client.Do(req)
	if err != nil {
		t.Error(err)
	}
}
