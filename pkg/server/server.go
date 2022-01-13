package server

import (
	"encoding/json"
	"net/http"

	"github.com/car-sales-website/pkg/services"
)

type Server struct {
	config        *Config
	personService *services.CarService
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/people", s.people)

	return mux
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}

	httpServer.ListenAndServe()
}

func (s *Server) people(w http.ResponseWriter, r *http.Request) {
	people := s.personService.FindAll()
	bytes, _ := json.Marshal(people)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewServer(config *Config, service *services.CarService) *Server {
	return &Server{
		config:        config,
		personService: service,
	}
}
