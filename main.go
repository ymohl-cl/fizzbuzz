package main

import (
	"flag"
	"os"

	"github.com/ymohl-cl/fizzbuzz/api"
	"github.com/ymohl-cl/gopkg/server"
)

var appName = flag.String("appName", "fizzbuzz", "your application name")
var help = flag.Bool("help", false, "show app configuration")

func init() {
	flag.Parse()
	if *help {
		if err := usage(); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

func main() {
	s, err := server.New(*appName)
	if err != nil {
		panic(err)
	}
	a, err := api.Init(*appName, s.SubRouter("/"))
	if err != nil {
		panic(err)
	}
	defer a.Close()
	if err = s.Start(); err != nil {
		panic(err)
	}
}
