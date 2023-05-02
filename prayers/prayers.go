package prayers

import (
	"fmt"
	"time"
)

type Prayer struct {
	Language     string `json:"Lang"`
	Creed        string `json:"Creed"`
	OurFather    string `json:"OurFather"`
	GloryBe      string `json:"GloryBe"`
	FatimaPrayer string `json:"FatimaPrayer"`
	FirstDecade  string `json:"FirstDecade"`
	SecondDecade string `json:"SecondDecade"`
	ThirdDecade  string `json:"ThirdDecade"`
	FourthDecade string `json:"FourthDecade"`
	FifthDecade  string `json:"FifthDecade"`
	SalveRegina  string `json:"SalveRegina"`
}

type Rosary struct {
	Weekday  string   `json:"Weekday"`
	Language string   `json:"Lang"`
	Decades  int      `json:"Decades"`
	Prayers  []Prayer `json:"Prayers"`
}



func getWeekday() string {
	return fmt.Sprint(time.Now().Weekday())
}
