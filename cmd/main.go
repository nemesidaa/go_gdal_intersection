package main

import (
	"fmt"
	"gogdal/internal/config"
	"gogdal/internal/http"
	"os"
)

var (
	cpath string
)

func parse() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: app <config-path>")
		os.Exit(2)
	}
	args := os.Args[1:]
	cpath = args[0]

	if cpath == "" {
		fmt.Fprintln(os.Stderr, "Config path cannot be empty")
		os.Exit(2)
	}
}

// @title Gogdal polygon intersection
// @version 1.0
// @description This is a sample API. By default, provides 8086 port as main. Provides to intersect multiple value of polygons in WKT or GeoJSON formats
// @host localhost:8086
// @BasePath /
func main() {
	parse()

	conf := new(config.Config)
	if err := conf.Parse(cpath); err != nil {
		panic(err)
	}
	serv, err := http.NewServer(conf)
	if err != nil {
		panic(err)
	}
	if err := serv.Serve(conf.Addr); err != nil {
		panic(err)
	}

}
