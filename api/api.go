package api

import (
	"database/sql"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	_ "github.com/lib/pq" // imported to postgres service
	"github.com/ymohl-cl/fizzbuzz/algo"
)

const (
	routeFizzBuzz = "fizzbuzz"
	routeStats    = "stats"
)

// API describe the endpoints and drivers needed
type API struct {
	driver *sql.DB
}

// Init get the configuration to the api
// Get the postgres driver and create the connexion
// Implement routes fizzbuzz and statistiques
func Init(appName string, router *echo.Group) (*API, error) {
	var err error
	var c Config
	var a API

	if err = envconfig.Process(appName, &c); err != nil {
		return nil, err
	}
	connSTR := c.connSTR()
	if a.driver, err = sql.Open("postgres", connSTR); err != nil {
		return nil, err
	}

	router.POST(routeFizzBuzz, a.FizzBuzz)
	router.GET(routeStats, a.Stats)
	return &a, nil
}

// Close the API
func (a API) Close() {
	a.driver.Close()
}

// FizzBuzz execute the fb algo with the parameters defines on the body
// Record the call if the request is success
func (a API) FizzBuzz(c echo.Context) error {
	var err error
	var data InputFB
	var result []string

	if err = c.Bind(&data); err != nil {
		c.Logger().Info(err.Error())
		return c.JSON(http.StatusBadRequest, InputFB{})
	}
	fb := algo.Algo{
		Start:  1,
		Limit:  data.Limit,
		N1:     data.Int1,
		N2:     data.Int2,
		Value1: data.Str1,
		Value2: data.Str2,
	}
	if result, err = fb.Run(); err != nil {
		c.Logger().Info(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err = a.record(data); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "please contact administrator")
	}
	return c.JSON(http.StatusOK, result)
}

// Stats return requests more performed
func (a API) Stats(c echo.Context) error {
	var output []OutputStat
	var err error

	if output, err = a.maxRecord(); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "please contact administrator")
	}
	if len(output) == 0 {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, output)
}
