package handler

import (
	"NameIdentification/domain"
	"NameIdentification/pkg/util"
	"encoding/json"
	_ "errors"
	"github.com/labstack/echo"
	_ "log"
	"net/http"
)

func GetFolderHandler(c echo.Context) error {

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
