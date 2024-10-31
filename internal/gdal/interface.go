package gdal

import (
	"gogdal/internal/config"
	mock "gogdal/internal/gdal/structs/mock"
)

type Worker interface {
	IntersectPolygons(polys ...string) (float64, bool, error)
}

func NewWorker(conf *config.Config, workerType string) Worker {
	switch workerType {
	// case "ord":
	// 	return ord.NewGdalWorker(conf)
	default:
		return mock.NewMockGdalWorker(conf)
	}
}
