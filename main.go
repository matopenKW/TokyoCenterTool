package main

import (
	"Carfare/handler"
	"Carfare/pkg"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func main() {
	router := gin.Default()

	// CORS
	c, _ := ini.Load(pkg.CONFIG_PATH)
	corsConfig := c.Section("CORS")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{corsConfig.Key("approvalServerUrl").String()}
	router.Use(cors.New(config))

	router.GET("/getCarfare", handler.GetCarfare)
	router.GET("/update", handler.UpdateCarfare)
	router.POST("/delete", handler.DeleteCarfare)

	router.Run()
}
