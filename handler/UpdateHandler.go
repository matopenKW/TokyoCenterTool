package handler

import (
	"Carfare/domain"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

func UpdateCarfare(db *gorm.DB, ctx *gin.Context) (interface{}, error) {

	ctx.Request.ParseForm()
	form := ctx.Request.Form
	log.Println(form)

	carfare := &domain.Carfare{
		SeqNo:      cnvInt(form["SeqNo"][0]),
		UserId:     "kazu",
		Date:       form["Date"][0],
		StartPoint: form["StartPoint"][0],
		EndPoint:   form["EndPoint"][0],
		Price:      cnvInt(form["Price"][0]),
	}

	err := update(db, carfare)
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}

func update(db *gorm.DB, carfare *domain.Carfare) error {
	return db.Debug().Model(&carfare).Where("seq_no=?", carfare.SeqNo).Update(&carfare).Error

}

func cnvInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
