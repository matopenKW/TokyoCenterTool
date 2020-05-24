package handler

import (
	"Carfare/domain"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

func RegistCarfare(db *gorm.DB, ctx *gin.Context) (interface{}, error) {

	ctx.Request.ParseForm()
	form := ctx.Request.Form
	log.Println(form)

	// userId := "kazu"

	// maxSeqNo, err := getMaxSeqNo(db, userId)
	// if err != nil {
	// 	return nil, err
	// }

	carfare := &domain.Carfare{
		SeqNo:      1,
		UserId:     "kazu",
		Date:       getFormData(form["Date"]),
		StartPoint: getFormData(form["StartPoint"]),
		EndPoint:   getFormData(form["EndPoint"]),
		Price:      cnvInt(getFormData(form["Price"])),
	}

	err := regist(db, carfare)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func getMaxSeqNo(db *gorm.DB, userId string) (int, error) {
	carfare := &domain.Carfare{}
	err := db.Debug().Model(&carfare).Last(&carfare, "user_id=?", userId).Error

	if err != nil {
		return 0, err
	}

	return carfare.SeqNo, nil
}

func regist(db *gorm.DB, carfare *domain.Carfare) error {
	return db.Debug().Updates(&carfare).Error
}

func getFormData(formData []string) string {
	if formData == nil || len(formData) < 1 {
		return ""
	} else {
		return formData[0]
	}
}
