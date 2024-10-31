package http

import (
	"encoding/json"
	"fmt"
	"gogdal/internal/config"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router     chi.Router
	Controller *Controller
	conf       *config.Config
}

func NewServer(conf *config.Config) *Server {
	serv := new(Server)
	serv.Controller = NewController(conf)
	serv.Router = chi.NewRouter()
	serv.Router.Get("/intersect_polygons", serv.IntersectPolygons)
	return serv
}

// ? Addr in format: 0.0.0.0:00000
func (s *Server) Serve(addr string) error {
	return http.ListenAndServe(addr, s.Router)
}

func (s *Server) IntersectPolygons(w http.ResponseWriter, r *http.Request) {
	var polys []string
	err := json.NewDecoder(r.Body).Decode(&polys)
	if err != nil {
		http.Error(w, "failed to read json data", http.StatusBadRequest)
		return
	}
	res, ok, err := s.Controller.IntersectPolygons()
	if !ok || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("%f", res))); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
