package main

import (
	"Carfare/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/getCarfare", handler.GetCarfare)

	router.Run()
}
