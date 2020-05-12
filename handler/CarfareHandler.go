package handler

import (
	"Carfare/pkg/util"

	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type Carfare struct {
	SeqNo      int    "json:seqNo"
	Date       string "json:date"
	StartPoint string "json:start_point"
	EndPoint   string "json:end_point"
	Price      int    "json:price"
}

func GetCarfare(ctx *gin.Context) {
	// list, err := getJSON()
	list, err := getData()
	if err != nil {
		log.Println("Error:", err)
		ret := `"error": "error です"`
		ctx.JSON(http.StatusInternalServerError, ret)
	} else {
		ctx.JSON(http.StatusOK, list)
	}
}

func getData() ([]*Carfare, error) {

	list, err := selectCarfare()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func selectCarfare() ([]*Carfare, error) {

	sqlCon, err := util.GetConnection()
	if err != nil {
		return nil, err
	}

	list := []*Carfare{}
	err = sqlCon.Table("carfare").Find(&list, "user_id=?", "kazu").Error

	return list, nil
}

func getJSON() ([]*Carfare, error) {

	bytes, err := util.ReadFile("json/carfareList.json")
	if err != nil {
		return nil, err
	}

	list := make([]*Carfare, 0, 0)
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
