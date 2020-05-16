package handler

import (
	"Carfare/domain"
	"Carfare/pkg/util"

	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func GetCarfare(ctx *gin.Context) {
	list, err := getJSON()
	//list, err := getData()
	if err != nil {
		log.Println("Error:", err)
		ret := `"error": "error です"`
		ctx.JSON(http.StatusInternalServerError, ret)
	} else {
		ctx.JSON(http.StatusOK, list)
	}
}

func getData() ([]*domain.Carfare, error) {

	list, err := selectCarfare()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func selectCarfare() ([]*domain.Carfare, error) {

	sqlCon, err := util.GetConnection()
	if err != nil {
		return nil, err
	}

	list := []*domain.Carfare{}
	err = sqlCon.Table("carfare").Find(&list, "user_id=?", "kazu").Error

	return list, nil
}

func getJSON() ([]*domain.Carfare, error) {

	bytes, err := util.ReadFile("json/carfareList.json")
	if err != nil {
		return nil, err
	}

	list := make([]*domain.Carfare, 0, 0)
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
