package http

import (
	"fmt"
	"gogdal/internal/gdal"

	"gogdal/internal/config"
)

type Controller struct {
	gdalw gdal.Worker
}

func NewController(conf *config.Config) (*Controller, error) {
	gdalWorker, err := gdal.NewWorker(conf, conf.WorkerType)
	if err != nil {
		return nil, fmt.Errorf("failed to create gdal worker: %w", err)
	}
	return &Controller{
		gdalw: gdalWorker,
	}, nil
}

func (ctrl *Controller) IntersectPolygons(polys ...string) (float64, bool, error) {
	return ctrl.gdalw.IntersectPolygons(polys...)
}
