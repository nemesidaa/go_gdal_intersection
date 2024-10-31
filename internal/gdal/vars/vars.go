package gdal_vars

import "regexp"

var (
	WktRegex     = regexp.MustCompile(`^\s*(POINT|LINESTRING|POLYGON|MULTIPOINT|MULTILINESTRING|MULTIPOLYGON|GEOMETRYCOLLECTION)\s*\(`)
	GeoJSONRegex = regexp.MustCompile(`^\s*\{.*"type"\s*:\s*"(Point|LineString|Polygon|MultiPoint|MultiLineString|MultiPolygon|GeometryCollection)".*\}`)
	GetType      = func(s string) GeometryType {
		if WktRegex.MatchString(s) {
			return WKT
		}
		if GeoJSONRegex.MatchString(s) {
			return GeoJSON
		}
		return Unknown
	}
)

type GeometryType byte

const (
	Unknown GeometryType = iota
	WKT
	GeoJSON
)
