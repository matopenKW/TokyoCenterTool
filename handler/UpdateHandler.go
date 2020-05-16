package handler

import (
	"Carfare/domain"
	"Carfare/pkg/util"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func UpdateCarfare(ctx *gin.Context) {

	ctx.Request.ParseForm()
	form := ctx.Request.Form
	log.Println(form)

	ret := make([]*domain.Carfare, 0, 0)
	carfare := &domain.Carfare{
		SeqNo:      1,
		Date:       form["Date"][0],
		StartPoint: form["StartPoint"][0],
		EndPoint:   form["EndPoint"][0],
		Price:      3000,
		//Price:      form["Price"][i],
	}

	ret = append(ret, carfare)

	err := update(carfare)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, `"error": "error です"`)
	} else {
		ctx.JSON(http.StatusOK, ret)
	}
}

func update(carfare *domain.Carfare) error {
	sqlCon, err := util.GetConnection()
	if err != nil {
		return err
	}

	err = sqlCon.Table("carfare").Update(carfare).Error

	return nil
}
