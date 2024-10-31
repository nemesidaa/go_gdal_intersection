package http

import (
	"gogdal/internal/gdal"

	"gogdal/internal/config"
)

type Controller struct {
	gdalw gdal.Worker
}

func NewController(conf *config.Config) *Controller {
	return &Controller{
		gdalw: gdal.NewWorker(conf, conf.WorkerType),
	}
}

func (c *Controller) IntersectPolygons(polys ...string) (float64, bool, error) {
	return c.gdalw.IntersectPolygons(polys...)
}
