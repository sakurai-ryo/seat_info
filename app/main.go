package main

import (
	"seat_info/controller"
	"time"

	"github.com/gin-gonic/gin"
)

const location = "Asia/Tokyo"

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.LoadHTMLGlob("./templates/**/*")
	engine.Static("/static", "./static")
	engine.GET("/", func(c *gin.Context) {
		controller.Controller(c)
	})
	engine.Run(":80")
}
