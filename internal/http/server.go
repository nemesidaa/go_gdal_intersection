package http

import (
	"encoding/json"
	"fmt"
	"gogdal/internal/config"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"

	_ "gogdal/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	Router     chi.Router
	Controller *Controller
	Logger     zerolog.Logger
	conf       *config.Config
}

func NewServer(conf *config.Config) (*Server, error) {
	serv := new(Server)
	var err error
	serv.Controller, err = NewController(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to create controller: %w", err)
	}
	serv.conf = conf
	// ! Z Я тут говна написал, но в целом понятно всё
	logdir := filepath.Dir(conf.Logdest)
	if err := os.MkdirAll(logdir, 0777); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}
	logf, err := os.OpenFile(conf.Logdest, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}
	level, err := zerolog.ParseLevel(conf.Loglevel)
	if err != nil {
		return nil, fmt.Errorf("failed to parse log level: %w", err)
	}
	serv.Logger = zerolog.New(logf).Level(level).With().Timestamp().Logger()
	serv.Router = chi.NewRouter()
	serv.Router.Use(serv.Log)
	serv.Router.Post("/intersect_polygons", serv.IntersectPolygons)
	if conf.Verbose == "true" {
		serv.Router.Get("/swagger/*", httpSwagger.WrapHandler)
	}
	serv.Logger.Debug().Msg("Server initialized")
	return serv, nil
}

// ? Addr in format: 0.0.0.0:00000
func (s *Server) Serve(addr string) error {
	s.Logger.Info().Str("addr", addr).Msg("starting server")
	s.Logger.Debug().Msg("Server is about to start listening")

	return http.ListenAndServe(addr, s.Router)
}

// intersect_polygons godoc
// @Summary Intersect polygons
// @Description Get a result of intersection of polygons
// @Accept json
// @Produce text/plain
// @Success 200 {number} float64 "1.0"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /intersect_polygons [post]
//
//	@Param polys body []string true "Polygons in WKT or GeoJSON format"
func (s *Server) IntersectPolygons(w http.ResponseWriter, r *http.Request) {
	s.Logger.Debug().Msg("Handling intersect polygons request")
	var polys []string
	err := json.NewDecoder(r.Body).Decode(&polys)
	if err != nil {
		http.Error(w, "failed to read json data", http.StatusBadRequest)
		s.Logger.Error().Err(err).Msg("failed to read json data")
		return
	}
	s.Logger.Debug().Msgf("Decoded polygons: %v", polys)
	res, ok, err := s.Controller.IntersectPolygons(polys...)
	if err != nil || !ok {
		w.WriteHeader(http.StatusBadRequest)
		s.Logger.Error().Err(err).Msg("failed to intersect polygons")
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("%f", res))); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.Logger.Error().Err(err).Msg("failed to write response")
		return
	}
	s.Logger.Debug().Msgf("Intersection result: %f", res)
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (s *Server) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Debug().Msg("Logging request")
		ww := &responseWriter{w, http.StatusOK}
		start := time.Now()
		next.ServeHTTP(ww, r)
		end := time.Now()
		s.Logger.Info().Fields(map[string]interface{}{
			"method":   r.Method,
			"path":     r.URL.Path,
			"status":   ww.status,
			"duration": end.Sub(start),
		}).Msg("")
		s.Logger.Debug().Msg("Request logged")
	})
}
