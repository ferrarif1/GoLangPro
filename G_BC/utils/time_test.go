package utils

import (
	"fmt"
	"testing"
)

func TestT1(t *testing.T) {

	now := GetCurrnetTime()
	fmt.Println("now: ", now)
	fmt.Println("nowUTC: ", now.UTC())
}

func TestT2(t *testing.T) {

	beginTime := GetCurrnetDayBegin()
	fmt.Println("beginTime: ", beginTime)
}

func TestT3(t *testing.T) {

	day1 := GetCurrnetTime()
	day2 := GetCurrnetTime().UTC()

	fmt.Println(DateAndClockEqual(day1, day2))
}
