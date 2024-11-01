package main

import (
	"gogdal/internal/config"
	"gogdal/internal/http"
	"os"
)

var (
	cpath string
)

func parse() {
	args := os.Args[1:]
	cpath = args[0]
	// * fmt.Println(os.Args)
}
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
