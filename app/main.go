package main

import (
	"seat_info/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// today := time.Date(2021, 2, 14, 12, 5, 6, 0, time.Local)
	engine.LoadHTMLGlob("./templates/**/*")
	engine.Static("/static", "./static")
	// today := time.Now()
	engine.GET("/", func(c *gin.Context) {
		controller.Controller(c)
	})
	engine.Run(":8080")
}
