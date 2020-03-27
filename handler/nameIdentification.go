package handler

import (
	"NameIdentification/domain"
	"NameIdentification/pkg/util"
	"encoding/json"
	_ "errors"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func NameIdentificationHandler(c echo.Context) error {

	teamList := make([]*domain.TeamName, 0, 0)
	log.Println(teamList)

	bytes, err := util.ReadFile("json/folder.json")
	if err != nil {
		return err
	}

	list := make([]*domain.Folder, 0, 0)
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, list)
}

func nameIdentification(list1 []*domain.TeamName, list2 []*domain.TeamName) ([]*domain.NameIdentificationFile, error) {

	return nil, nil
}
