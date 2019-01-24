package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/ymohl-cl/gopkg/httput"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const (
	uri = "http://localhost:4242/"
)

func TestFizzBuzz(t *testing.T) {
	// should return an error bicause data can't be parse by echo.Bind
	func() {
		// init
		expectedResponse := `{"int1":0,"int2":0,"limit":0,"str1":"","str2":""}`
		req := httptest.NewRequest(http.MethodPost, uri+routeFizzBuzz, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		ctx := httput.NewContext(req)
		a := API{}

		// assert
		if assert.NoError(t, a.FizzBuzz(ctx.Input)) {
			assert.Equal(t, http.StatusBadRequest, ctx.Rec.Code)
			assert.Equal(t, expectedResponse, ctx.Rec.Body.String())
		}
	}()

	// should return an error because fizzbuzz algo return an error
	func() {
		// init
		expectedResponse := `"start must be smaller than limit"`
		data := InputFB{Limit: -10, Int1: 3, Int2: 5, Str1: "fizz", Str2: "buzz"}
		req, err := httput.RequestJSON(http.MethodPost, uri+routeFizzBuzz, &data)
		assert.NoError(t, err)
		ctx := httput.NewContext(req)
		a := API{}

		// assert
		if assert.NoError(t, a.FizzBuzz(ctx.Input)) {
			assert.Equal(t, http.StatusBadRequest, ctx.Rec.Code)
			assert.Equal(t, expectedResponse, ctx.Rec.Body.String())
		}
	}()

	// should return an error bicause record on postgres failed
	func() {
		// init
		db, mock, err := sqlmock.New()
		defer db.Close()
		assert.NoError(t, err)
		data := InputFB{Limit: 2, Int1: 1, Int2: 2, Str1: "f", Str2: "b"}
		req, err := httput.RequestJSON(http.MethodPost, uri+routeFizzBuzz, &data)
		assert.NoError(t, err)
		ctx := httput.NewContext(req)
		mock.ExpectExec("INSERT INTO stats").WithArgs("212fb", 2, 1, 2, "f", "b", 1).WillReturnError(fmt.Errorf("some error"))
		a := API{driver: db}

		// assert
		if assert.NoError(t, a.FizzBuzz(ctx.Input)) {
			assert.Equal(t, http.StatusInternalServerError, ctx.Rec.Code)
			assert.Equal(t, `"please contact administrator"`, ctx.Rec.Body.String())
		}
	}()

	// default: should be ok
	func() {
		// init
		expectedResult := `["f","fb","f"]`
		db, mock, err := sqlmock.New()
		defer db.Close()
		assert.NoError(t, err)
		data := InputFB{Limit: 3, Int1: 1, Int2: 2, Str1: "f", Str2: "b"}
		req, err := httput.RequestJSON(http.MethodPost, uri+routeFizzBuzz, &data)
		assert.NoError(t, err)
		ctx := httput.NewContext(req)
		mock.ExpectExec("INSERT INTO stats").WithArgs("312fb", 3, 1, 2, "f", "b", 1).WillReturnResult(sqlmock.NewResult(1, 1))
		a := API{driver: db}

		// assert
		if assert.NoError(t, a.FizzBuzz(ctx.Input)) {
			assert.Equal(t, http.StatusOK, ctx.Rec.Code)
			assert.Equal(t, expectedResult, ctx.Rec.Body.String())
		}
	}()
}

func TestStats(t *testing.T) {
	// should return an error because postgres failed
	func() {
		// init
		db, mock, err := sqlmock.New()
		defer db.Close()
		assert.NoError(t, err)
		data := InputFB{Limit: 3, Int1: 1, Int2: 2, Str1: "f", Str2: "b"}
		req, err := httput.RequestJSON(http.MethodGet, uri+routeStats, &data)
		assert.NoError(t, err)
		ctx := httput.NewContext(req)
		mock.ExpectExec("SELECT (.+)").WithArgs().WillReturnError(fmt.Errorf("some error"))
		a := API{driver: db}

		// assert
		if assert.NoError(t, a.Stats(ctx.Input)) {
			assert.Equal(t, http.StatusInternalServerError, ctx.Rec.Code)
			assert.Equal(t, `"please contact administrator"`, ctx.Rec.Body.String())
		}
	}()

	// default: should be ok without content
	func() {
		// init
		db, mock, err := sqlmock.New()
		defer db.Close()
		assert.NoError(t, err)
		data := InputFB{Limit: 3, Int1: 1, Int2: 2, Str1: "f", Str2: "b"}
		req, err := httput.RequestJSON(http.MethodGet, uri+routeStats, &data)
		assert.NoError(t, err)
		ctx := httput.NewContext(req)
		columns := []string{"limit_range", "int1", "int2", "str1", "str2", "nb_record"}
		mock.ExpectQuery("SELECT (.+)").WithArgs().WillReturnRows(sqlmock.NewRows(columns))
		a := API{driver: db}

		// assert
		if assert.NoError(t, a.Stats(ctx.Input)) {
			assert.Equal(t, http.StatusNoContent, ctx.Rec.Code)
			assert.Equal(t, "", ctx.Rec.Body.String())
		}
	}()

	// default: should be ok with content
	func() {
		// init
		expectedResult := `[{"int1":1,"int2":2,"limit":3,"str1":"f","str2":"b","count":3}]`
		db, mock, err := sqlmock.New()
		defer db.Close()
		assert.NoError(t, err)
		data := InputFB{Limit: 3, Int1: 1, Int2: 2, Str1: "f", Str2: "b"}
		req, err := httput.RequestJSON(http.MethodGet, uri+routeStats, &data)
		assert.NoError(t, err)
		ctx := httput.NewContext(req)
		columns := []string{"limit_range", "int1", "int2", "str1", "str2", "nb_record"}
		mock.ExpectQuery("SELECT (.+)").WithArgs().WillReturnRows(sqlmock.NewRows(columns).AddRow(3, 1, 2, "f", "b", 3))
		a := API{driver: db}

		// assert
		if assert.NoError(t, a.Stats(ctx.Input)) {
			assert.Equal(t, http.StatusOK, ctx.Rec.Code)
			assert.Equal(t, expectedResult, ctx.Rec.Body.String())
		}
	}()
}
