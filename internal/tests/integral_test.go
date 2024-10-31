package tests

import (
	"gogdal/internal/config"
	gdal_ord "gogdal/internal/gdal/structs/ord"
	"testing"
)

func TestIntegralIntersection(t *testing.T) {
	worker, err := gdal_ord.NewGdalWorker(&config.Config{Spatref: 4326})
	if err != nil {
		t.Fatalf("failed to create worker: %v", err)
	}
	area, ok, err := worker.IntersectPolygons(Json, Wkt)
	if err != nil {
		t.Fatalf("failed to intersect polygons: %v", err)
	}
	if !ok {
		t.Fatal("polygons do not intersect")
	}
	if area != 1.0 {
		t.Fatalf("expected area to be 1.0, got %f", area)
	}
}
