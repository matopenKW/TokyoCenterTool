package handler

import (
	"Carfare/domain"
	"Carfare/pkg/util"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func DeleteCarfare(ctx *gin.Context) {

	ctx.Request.ParseForm()
	form := ctx.Request.Form
	log.Println(form)

	ret := make([]*domain.Carfare, 0, 0)
	for i, _ := range form["SeqNo"] {
		carfare := &domain.Carfare{
			SeqNo:      1,
			Date:       form["Date"][i],
			StartPoint: form["StartPoint"][i],
			EndPoint:   form["EndPoint"][i],
			Price:      3000,
			//Price:      form["Price"][i],
		}

		ret = append(ret, carfare)
	}

	err := delete(1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, `"error": "error です"`)
	} else {
		ctx.JSON(http.StatusOK, ret)
	}
}

func delete(seqNo int) error {
	sqlCon, err := util.GetConnection()
	if err != nil {
		return err
	}

	err = sqlCon.Table("carfare").Delete("seq_no=?", seqNo).Error
	return err
}
