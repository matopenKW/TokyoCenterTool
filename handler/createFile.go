package handler

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

const (
	credentialFilePath = "pkg/conf/credentials.json"
)

func CreateFileHandler(c echo.Context) error {
	client, err := getClient(credentialFilePath)
	if err != nil {
		return err
	}
	driveService, _ := drive.New(client)
	folder := &drive.File{}
	folder.Name = "testFolder"
	folder.MimeType = "application/vnd.google-apps.folder"
	res, err := driveService.Files.Create(folder).Do()
	if err != nil {
		return err
	}
	log.Println(res)

	return c.String(http.StatusOK, "OK")
}

func getClient(credentialFilePath string) (*http.Client, error) {
	credentialInfo, err := ioutil.ReadFile(credentialFilePath)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON(credentialInfo, drive.DriveFileScope)
	if err != nil {
		return nil, err
	}

	return config.Client(oauth2.NoContext), nil
}
