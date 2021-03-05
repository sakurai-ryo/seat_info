package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"seat_info/shared"
	"time"

	"github.com/gin-gonic/gin"
)

func Controller(c *gin.Context) {
	today := time.Now()
	isOpen := shared.IsOpen(today)
	// logging(c)

	if isOpen {
		c.HTML(http.StatusOK, "open/open.html", gin.H{
			"color": "color: red;",
		})
	} else {
		c.HTML(http.StatusOK, "close/close.html", gin.H{
			"color": "color: red;",
		})
	}
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
