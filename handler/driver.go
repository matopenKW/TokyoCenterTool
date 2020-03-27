package handler

import (
	"NameIdentification/domain"
	"NameIdentification/pkg"
	"errors"
	"github.com/labstack/echo"
	_ "log"
	"net/http"
)

type DriveList struct {
	Exsample []*domain.Exsample
}

func DriveHandler(c echo.Context) error {
	const FOLDER_NAME = "DriveTest"

	service, err := pkg.GetService()
	if err != nil {
		return err
	}

	files, err := pkg.GetFiles(service, FOLDER_NAME)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return errors.New("No files in " + FOLDER_NAME)
	}

	driveList := make([]*domain.Exsample, 0, 0)

	for _, i := range files {
		exsample := &domain.Exsample{
			Id:   i.Id,
			Name: i.Name,
		}

		driveList = append(driveList, exsample)
	}

	res := &DriveList{
		Exsample: driveList,
	}

	return c.Render(http.StatusOK, "drive.html", res)
}
