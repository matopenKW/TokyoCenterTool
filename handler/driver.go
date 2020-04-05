package handler

import (
	"NameIdentification/domain"
	"NameIdentification/pkg"
	"errors"
	"github.com/labstack/echo"
	_ "io/ioutil"
	"log"
	"net/http"
	"os"
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
		log.Println(err)
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

		req := service.Files.Get(i.Id)
		output, err := os.Create("docs/file2.xlsx")
		if err != nil {
			return err
		}
		defer output.Close()

		res, err := req.Download()
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(res)

		//res := service.Files.Export(i.Id, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

		//log.Println(res)

		// url := "https://www.googleapis.com/drive/v3/files/19DKN2RCBKBjSEaTGpDxuT9DuVHkjuhUttUvfuHBvwfM?alt=media"

		// resp, _ := http.Get(url)
		// defer resp.Body.Close()

		// byteArray, _ := ioutil.ReadAll(resp.Body)
		// log.Println(string(byteArray))
		log.Println("success")

		driveList = append(driveList, folder)
	}

	return c.JSON(http.StatusOK, driveList)
}
