package main

import (
	"log"
	"seat_info/controller"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// today := time.Date(2021, 2, 14, 12, 5, 6, 0, time.Local)
	engine.LoadHTMLGlob("./templates/**/*")
	engine.Static("/static", "./static")
	// today := time.Now()
	engine.GET("/", func(c *gin.Context) {
		log.Print(time.Now())
		controller.Controller(c)
	})
	engine.Run(":80")
}
