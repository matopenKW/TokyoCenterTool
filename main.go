package main

import (
	"NameIdentification/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error= ${error}\n",
	}))
	e.Use(middleware.Recover())

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.GET("/hello", handler.HelloHandler)
	e.GET("/testDrive", handler.DriveHandler)
	e.GET("/testDrive2", handler.DriveHandler2)
	e.GET("/getFolder", handler.GetFolderHandler)
	e.GET("/nameIdentification", handler.NameIdentificationHandler)
	e.GET("/createFile", handler.CreateFileHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
