package main

import (
	"encoding/json"
	"log"
	"net/http"
	"seat_info/shared"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// today := time.Date(2021, 2, 14, 12, 5, 6, 0, time.Local)
	engine.LoadHTMLGlob("./templates/**/*")
	today := time.Now()
	engine.GET("/", func(c *gin.Context) {
		// logging(c)

		var path string
		if shared.IsOpen(today) {
			path = "open/open.html"
		} else {
			path = "close/close.html"
		}
		c.HTML(http.StatusOK, path, gin.H{
			"message": "hello World",
		})
	})
	engine.Run(":8080")
}

func logging(c *gin.Context) {
	if res, err := json.Marshal(&c.Request.Header); err != nil {
		log.Print(err)
	} else {
		log.Print(string(res))
	}

	if res, err := json.Marshal(&c.Request.URL); err != nil {
		log.Print(err)
	} else {
		log.Print(string(res))
	}

	if res, err := json.Marshal(&c.Request.Body); err != nil {
		log.Print(err)
	} else {
		log.Print(string(res))
	}
}
