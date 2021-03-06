package controller

import (
	"encoding/json"
	"fmt"
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
		// seat, err := Request()
		seat := []string{"a4" ,"b3","a7" ,"a2","b1" ,"1" ,"4" ,"a6" ,"a8" ,"D1" ,"10" ,"F1" ,"7" ,"b2" "a1" ,"a3" ,"C1" ,"8" ,"3" ,"E1"}
		if err != nil {
			log.Fatal(err)
		}
		log.Print(seat)
		c.HTML(http.StatusOK, "open/open.html", gin.H{
			"color": "color: red;",
		})
	} else {
		c.HTML(http.StatusOK, "close/close.html", gin.H{
			"color": "color: red;",
		})
	}
}

func createHTML(id string, color string) string {
	tem := fmt.Sprintf(`<div id= "%s" class="wrapper counter col-1" style="backgroundColor: coral">
    <div class="content">1人</div>
</div>`, id)
	return tem
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
