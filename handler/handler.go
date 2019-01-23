package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/fizzbuzz"
)

const (
	routeFizzBuzz = "fizzbuzz"
	routeStats    = "statistique"
)

type Handler struct {
	// bdd
}

func Init(router *echo.Group) error {
	h := Handler{}
	router.Use(metric.Record())
	// save metric and aggregate it how ?
	// hash the request and save the hash and the numbe time to used
	// warn to json order keys
	router.PUT(routeFizzBuzz, h.FizzBuzz)
	router.GET(routeStats, h.Stats)
	return nil
}

func (h Handler) FizzBuzz(c echo.Context) error {
	var err error
	var data Input

	if err = c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, Input{})
	}
	return c.JSON(http.StatusOK, fizzbuzz.Run(data.n1, data.n2, data.v1, data.v2))
}

func (h Handler) Stats(c echo.Context) error {
	var err error
	var data Input

	// TODO:
	// create a new object which embeded data to prove th enumber time this endpoint called with
	// this parameters
	if data, err = GetBestCalled(routeFizzBuzz); err != nil {
		return c.JSON(http.StatusInternalServerError, "can't get the best called parameter")
	}
	return c.JSON(http.StatusOK, data)
}
