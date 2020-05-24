package main

import (
	"Carfare/handler"
	"Carfare/pkg"
	"Carfare/pkg/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	// CORS
	c, _ := ini.Load(pkg.CONFIG_PATH)
	corsConfig := c.Section("CORS")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{corsConfig.Key("approvalServerUrl").String()}
	router.Use(cors.New(config))

	router.GET("/getCarfare", getResponce(handler.GetCarfare))
	router.GET("/regist", getResponce(handler.RegistCarfare))
	router.GET("/update", getResponce(handler.UpdateCarfare))
	router.POST("/delete", handler.DeleteCarfare)

	router.Run()
}

func getResponce(handler func(db *gorm.DB, ctx *gin.Context) (interface{}, error)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		db, err := util.GetConnection()
		defer db.Close()
		if err != nil {
			log.Println("Error:", err)
			ctx.JSON(http.StatusOK, `"error": "error です"`)
		}

		// Transaction Start
		tx := db.Begin()
		defer tx.Rollback()

		ret, err := handler(db, ctx)
		if err != nil {

			log.Println("Error:", err)
			ctx.JSON(http.StatusOK, `"error": "error です"`)
		} else {
			tx.Commit()

			ctx.JSON(http.StatusOK, ret)
		}
	}
}
