package endpoint

import (
	"encoding/json"
	"log"
	"net/http"
)

type Service interface {
	GetSheetsData(sId string, rng string) []string
	InitCache()
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) HandleClientRequest(w http.ResponseWriter, r *http.Request) {
	e.s.InitCache()
	p, err := json.Marshal(e.s.GetSheetsData(r.URL.Query().Get("s_id"), r.URL.Query().Get("range")))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(p)
}
