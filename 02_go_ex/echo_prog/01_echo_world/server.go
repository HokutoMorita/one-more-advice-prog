package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/item/:id", func(c echo.Context) error {
		paramValue := c.Param("id")
		formValue := c.FormValue("name")
		return c.String(http.StatusOK, "formValue = "+formValue + " paramValue = "+paramValue)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
