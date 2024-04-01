package helpers

import (
	"fmt"
	"strings"
	"time"
)

type ValidTime struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func StringToTimeDuration(str string) time.Duration {
	var n int64
	fmt.Sscan(str, &n)
	return time.Duration(n)
}

//hh:mm:ss
func ParseTimeValidate(timeStr string) time.Duration {
	timer := strings.Split(timeStr, ":")
	if len(timer) != 3 {
		return time.Microsecond
	}
	return StringToTimeDuration(timer[0])*time.Hour +
		StringToTimeDuration(timer[1])*time.Minute +
		StringToTimeDuration(timer[2])*time.Second
}

func ValidateTime(validTimes []ValidTime) bool {
	//init the loc
	loc, _ := time.LoadLocation("Asia/Bangkok")

	//set timezone,
	now := time.Now().In(loc)
	currentDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	for _, validTime := range validTimes {
		minTime := currentDate.Add(ParseTimeValidate(validTime.From))
		maxTime := currentDate.Add(ParseTimeValidate(validTime.To))
		if now.After(minTime) && now.Before(maxTime) {
			return true
		}
	}

	return false
}
