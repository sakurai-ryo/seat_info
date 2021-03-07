package main

import (
	"log"
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
