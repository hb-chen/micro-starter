package router

import (
	"github.com/labstack/echo"
	"net/http"
)

func NewRouter() *echo.Echo {
	// Echo instance
	e := echo.New()

	e.GET("/", handler())
	e.GET("/handler1", handler())
	e.GET("/handler2", handler())

	return e
}

func handler() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := map[string]interface{}{
			"RequestURI": c.Request().RequestURI,
		}
		c.JSON(http.StatusOK, resp)
		return nil
	}
}
