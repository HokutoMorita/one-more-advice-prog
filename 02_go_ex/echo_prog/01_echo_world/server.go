package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/item", func(c echo.Context) error {
		name := c.FormValue("name")
		return c.String(http.StatusOK, "create name = "+name)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
