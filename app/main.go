package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// engine.LoadHTMLGlob("templates/*.html")
	// today := time.Date(2021, 2, 14, 12, 5, 6, 0, time.Local)
	engine.GET("/", func(c *gin.Context) {
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

		// var path string
		// if shared.IsOpen(today) {
		// 	path = "open/index.html"
		// } else {
		// 	path = "close/index.html"
		// }
		// log.Print(path)
		// c.HTML(http.StatusOK, path, gin.H{
		// 	// htmlに渡す変数を定義
		// 	"message": "hello World",
		// })
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":80")
}
