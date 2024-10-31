package tests

import (
	gdal_vars "gogdal/internal/gdal/vars"
	"testing"
)

func TestDiffRegex(t *testing.T) {

	if gdal_vars.GeoJSONRegex.MatchString(Json) != gdal_vars.WktRegex.MatchString(Wkt) {
		t.Errorf("expected %s to match %s", Json, Wkt)
	}
}
