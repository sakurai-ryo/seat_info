package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"seat_info/shared"
	"time"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Seat     string
	Category string
}
type Results []Result

type NumOfCategory struct {
	One     int
	Table   int
	Sofa    int
	Counter int
}

// TODO: controllerに全部書いてるがあんまりよくない
// modelやらにいつか分ける
func Controller(c *gin.Context) {
	today := time.Now()
	isOpen := shared.IsOpen(today)
	// logging(c)

	if !isOpen {
		res, err := request()
		// res := []string{"a4", "b3", "a7", "a2", "b1", "1", "4", "a6", "a8", "D1", "10", "F1", "7", "b2", "a1", "a3", "C1", "8", "3", "E1"}
		if err != nil {
			log.Fatal(err)
		}
		// var res Results
		// res = Results{{"a4", "2F-a"}, {"C1", "2F"}, {"b4", "2F-b"}, {"a1", "2F-a"}, {"a7", "2F-a"}, {"5", "1F"}, {"3", "1F"}}

		num := calcNumOfSeat(res)

		c.HTML(http.StatusOK, "open/open.html", gin.H{
			"one":     10 - num.One,
			"table":   8 - num.Table,
			"sofa":    7 - num.Sofa,
			"counter": 4 - num.Counter,
			"seats":   getSeats(res),
		})
	} else {
		c.HTML(http.StatusOK, "close/close.html", gin.H{
			"color": "color: red;",
		})
	}
}

// TODO: 下で何回もループ回すのは無駄
func getSeats(re Results) []string {
	var seats []string
	for _, r := range re {
		seats = append(seats, r.Seat)
	}
	return seats
}

func calcNumOfSeat(re Results) NumOfCategory {
	var num NumOfCategory
	for _, r := range re {
		if r.Category == "1F" {
			num.One++
		}
		if r.Category == "2F-a" {
			num.Table++
		}
		if r.Category == "2F" {
			num.Sofa++
		}
		if r.Category == "2F-b" {
			num.Counter++
		}
	}
	return num
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
