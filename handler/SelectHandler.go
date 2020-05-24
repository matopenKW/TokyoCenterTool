package handler

import (
	"Carfare/domain"
	"Carfare/pkg/util"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "log"
)

func GetCarfare(db *gorm.DB, ctx *gin.Context) (interface{}, error) {

	isJson := false
	userId := "kazu"

	if isJson {
		return getJson()
	} else {
		return getData(userId)
	}
}

func getData(userId string) ([]*domain.Carfare, error) {

	sqlCon, err := util.GetConnection()
	if err != nil {
		return nil, err
	}

	list := []*domain.Carfare{}
	err = sqlCon.Model(&list).Find(&list, "user_id=?", userId).Error

	return list, nil
}

func getJson() ([]*domain.Carfare, error) {

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
