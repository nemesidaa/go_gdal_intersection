package gdal_ord

// import (
// 	"fmt"
// 	"gogdal/internal/config"
// 	vars "gogdal/internal/gdal/vars"

// 	"github.com/lukeroth/gdal"
// )

// // * Getting spatref from config
// type GdalWorker struct {
// 	spatref gdal.SpatialReference
// }

// func NewGdalWorker(conf *config.Config) *GdalWorker {
// 	spatialRef := gdal.CreateSpatialReference("")
// 	spatialRef.FromEPSG(conf.Spatref)
// 	return &GdalWorker{
// 		spatref: spatialRef,
// 	}
// }

// // ? Getting polys in WKT or GeoJSON formats, outputting square and existance of the intersection
// func (gdalw *GdalWorker) IntersectPolygons(polys ...string) (float64, bool, error) {
// 	ppolys := make([]gdal.Geometry, 0)
// 	for idx, upoly := range polys {
// 		var err error
// 		switch vars.GetType(upoly) {
// 		case vars.WKT:
// 			ppolys[idx], err = gdal.CreateFromWKT(upoly, gdalw.spatref)
// 			if err != nil {
// 				fmt.Printf("failed to create geometry from WKT: %w", err)
// 				continue
// 			}
// 		case vars.GeoJSON:
// 			ppolys[idx] = gdal.CreateFromJson(upoly)
// 		default:
// 			continue
// 		}

// 	}
// 	var resp gdal.Geometry
// 	for _, poly := range ppolys {
// 		resp = resp.Intersection(poly)
// 	}
// 	a := resp.Area()
// 	return a, a != 0, nil
// }
