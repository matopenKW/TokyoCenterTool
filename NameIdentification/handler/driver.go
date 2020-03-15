package handler

import (
	"NameIdentification/domain"
	"github.com/labstack/echo"
	"net/http"
)

func DriveHandler(c echo.Context) error {
	exsample := &domain.Exsample{
		Id:   "id",
		Name: "name",
	}
	return c.Render(http.StatusOK, "hello.html", exsample)
}
