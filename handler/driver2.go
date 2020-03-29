package handler

import (
	"github.com/labstack/echo"
	"io"
	"log"
	"net/http"
	"os"
)

func DriveHandler2(c echo.Context) error {
	file_id := "1m5wuPbfvbxHIadcIn3z7bgbDZGUFVMvH"
	url := "https://docs.google.com/uc?export=download&id=" + file_id
	fileName := "docs/file.xlsx"
	log.Println(url)
	log.Println("Downloading file...")

	output, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return err
	}
	defer response.Body.Close()

	log.Println(response)

	n, err := io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	log.Println(n, "bytes downloaded")

	return c.JSON(http.StatusOK, string(n))
}
