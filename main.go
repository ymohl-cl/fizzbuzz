package main

import (
	"github.com/ymohl-cl/gopkg/server"
)

func main() {
	s, err := server.New("fizzbuzz-api")
	if err != nil {
		panic(err)
	}
	sr := s.SubRouter("fizzbuzz")
	if err = handler.Init(sr); err != nil {
		panic(err)
	}
	if err = s.Start(); err != nil {
		panic(err)
	}
}
