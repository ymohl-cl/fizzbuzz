package main

import (
	"github.com/ymohl-cl/fizzbuzz/api"
	"github.com/ymohl-cl/gopkg/server"
)

func main() {
	s, err := server.New("fizzbuzz")
	if err != nil {
		panic(err)
	}
	a, err := api.Init("fizzbuzz", s.SubRouter("/"))
	if err != nil {
		panic(err)
	}
	defer a.Close()
	if err = s.Start(); err != nil {
		panic(err)
	}
}
