package tests

import (
	gdal_vars "gogdal/internal/gdal/vars"
	"testing"
)

func TestDiffRegex(t *testing.T) {
	json := `{"type":"Polygon","coordinates":[[100,0],[101,0],[101,1],[100,1],[100,0]]}`
	wkt := "POLYGON((100 0,101 0,101 1,100 1,100 0))"

	if gdal_vars.GeoJSONRegex.MatchString(json) != gdal_vars.WktRegex.MatchString(wkt) {
		t.Errorf("expected %s to match %s", json, wkt)
	}
}
