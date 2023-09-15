package util

import (
	"fmt"
	"time"
)

func LitCal() {
	// TODO implement API for liturgical information
	api := "http://calapi.inadiutorium.cz/api-doc"
	fmt.Println(api)
}

func DayOfWeek() string {
	return time.Now().Weekday().String()
}
