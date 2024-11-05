package gdal_ord

import (
	"fmt"
	"gogdal/internal/config"
	vars "gogdal/internal/gdal/vars"

	"github.com/airbusgeo/godal"
)

// * Getting spatref from config
type GdalWorker struct {
	spatref *godal.SpatialRef
}

func NewGdalWorker(conf *config.Config) (*GdalWorker, error) {
	spatialRef, err := godal.NewSpatialRefFromEPSG(conf.Spatref)
	if err != nil {
		return nil, fmt.Errorf("failed to create spatial reference: %w", err)
	}
	return &GdalWorker{
		spatref: spatialRef,
	}, nil
}

// ? Getting polys in WKT or GeoJSON formats, outputting square and existance of the intersection
func (gdalw *GdalWorker) IntersectPolygons(polys ...string) (float64, bool, error) {
	ppolys, err := gdalw.TranslateToPolygons(polys...)
	if err != nil {
		return 0, false, fmt.Errorf("failed to translate polygons: %w", err)
	}
	resp := ppolys[0]
	for _, poly := range ppolys[1:] {
		var err error
		resp, err = resp.Intersection(poly)
		if err != nil {
			return 0, false, fmt.Errorf("failed to intersect polygons: %w", err)
		}
	}
	a := resp.Area()
	return a, a != 0, nil
}

func (gdalw *GdalWorker) TranslateToPolygons(polys ...string) ([]*godal.Geometry, error) {
	ppolys := make([]*godal.Geometry, 0, len(polys))
	for _, upoly := range polys {
		switch vars.GetType(upoly) {
		case vars.WKT:
			ppoly, err := godal.NewGeometryFromWKT(upoly, gdalw.spatref)
			if err != nil || ppoly == nil {
				fmt.Printf("failed to create geometry from WKT: %v", err)
				continue
			}
			ppolys = append(ppolys, ppoly)
		case vars.GeoJSON:
			ppoly, err := godal.NewGeometryFromGeoJSON(upoly)
			if err != nil || ppoly == nil {
				continue
			}
			ppolys = append(ppolys, ppoly)
		default:
			continue
		}

	}
	return ppolys, nil
}
