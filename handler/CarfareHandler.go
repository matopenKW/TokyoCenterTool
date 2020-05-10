package handler

import (
	"Carfare/pkg/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Carfare struct {
	SeqNo int    "json:seqNo"
	Date  string "json:date"
	Start string "json:start"
	End   string "json:end"
	Price int    "json:price"
}

func GetCarfare(ctx *gin.Context) {
	list, err := getJSON()
	if err != nil {
		log.Println("Error:", err)
		ret := `"error": "error です"`
		ctx.JSON(http.StatusInternalServerError, ret)
	} else {
		ctx.JSON(http.StatusOK, list)
	}
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
