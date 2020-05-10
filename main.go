package main

import (
	"Carfare/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

const CONFIG_PATH = "pkg/conf/config.conf"

func main() {
	router := gin.Default()

	// CORS
	c, _ := ini.Load(CONFIG_PATH)
	corsConfig := c.Section("CORS")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{corsConfig.Key("approvalServerUrl").String()}
	router.Use(cors.New(config))

	router.GET("/getCarfare", handler.GetCarfare)

	router.Run()
}
