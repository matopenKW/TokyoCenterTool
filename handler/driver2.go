package handler

import (
	"github.com/PuerkitoBio/goquery"
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

	output2, err := os.Create("docs/file2.xlsx")
	if err != nil {
		return err
	}
	defer output2.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return err
	}
	defer response.Body.Close()

	log.Println(response)

	// n, err := io.Copy(output, response.Body)
	n, err := goquery.NewDocumentFromReader(
		io.TeeReader(response.Body, io.MultiWriter(output, output2)),
	)
	if err != nil {
		return err
	}

	log.Println(n, "bytes downloaded")

	return c.JSON(http.StatusOK, "")
}
