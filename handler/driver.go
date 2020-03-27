package handler

import (
	"NameIdentification/domain"
	"NameIdentification/pkg"
	"errors"
	"github.com/labstack/echo"
	_ "io/ioutil"
	"log"
	"net/http"
)

type DriveList struct {
	FolderList []*domain.Folder
}

func DriveHandler(c echo.Context) error {
	const FOLDER_NAME = "develop"

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

	driveList := make([]*domain.Folder, 0, 0)

	for _, i := range files {
		folder := &domain.Folder{
			Id:   i.Id,
			Name: i.Name,
		}

		res, err := service.Files.Get("1693YKeYqplM4iStd2-p63d-a97tZhByt").Download()
		if err != nil {
			return err
		}

		log.Println(res)

		// url := "https://www.googleapis.com/drive/v3/files/19DKN2RCBKBjSEaTGpDxuT9DuVHkjuhUttUvfuHBvwfM?alt=media"

		// resp, _ := http.Get(url)
		// defer resp.Body.Close()

		// byteArray, _ := ioutil.ReadAll(resp.Body)
		// log.Println(string(byteArray))

		driveList = append(driveList, folder)
	}

	return c.JSON(http.StatusOK, driveList)
}
