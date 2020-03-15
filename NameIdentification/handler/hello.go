package handler

import (
	"NameIdentification/domain"
	"github.com/labstack/echo"
	"net/http"
)

func HelloHandler(c echo.Context) error {
	exsample := &domain.Exsample{
		Id:   "id",
		Name: "name",
	}
	return c.Render(http.StatusOK, "hello.html", exsample)
}
