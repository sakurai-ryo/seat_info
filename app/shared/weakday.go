package shared

import (
	"log"
	"time"
)

// IsOpen ... if cafe is open, return true
func IsOpen(t time.Time) bool {
	log.Println("date: ", t.Weekday())
	if isThirdSunday(t) {
		return false
	}
	h := t.Hour()

	if isSaturday(t) || isSunday(t) {
		if 11 < h && h < 19 {
			return true
		}
		return false
	}
	if 8 < h && h < 19 {
		return true
	}
	return false
}

func isSaturday(t time.Time) bool {
	if t.Weekday() == 6 {
		return true
	}
	return false
}

func isSunday(t time.Time) bool {
	if t.Weekday() == 0 {
		return true
	}
	return false
}

// IsThirdSunday ... 第三日曜日か判定する
func isThirdSunday(t time.Time) bool {
	fd := time.Date(t.Year(), t.Month(), 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local).Weekday()
	sMap := map[time.Weekday]int{0: 22, 1: 21, 2: 20, 3: 19, 4: 18, 5: 17, 6: 18}
	d := t.Day()
	if sMap[fd] == d {
		return true
	}
	return false
}
