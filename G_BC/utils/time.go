package utils

import "time"

var TIME_LOCATION *time.Location

func init() {
	var err error
	TIME_LOCATION, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}

func GetCurrnetTime() time.Time {
	return time.Now().In(TIME_LOCATION)
}

func GetCurrnetDayBegin() time.Time {
	now := GetCurrnetTime()

	y, m, d := now.Date()

	return time.Date(y, m, d, 00, 00, 00, 00, TIME_LOCATION)
}

func DateEqual(day1, day2 time.Time) bool {
	day1 = day1.In(TIME_LOCATION)
	day2 = day2.In(TIME_LOCATION)

	y1, m1, d1 := day1.Date()
	y2, m2, d2 := day2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

func DateAndClockEqual(day1, day2 time.Time) bool {
	day1 = day1.In(TIME_LOCATION)
	day2 = day2.In(TIME_LOCATION)

	y1, m1, d1 := day1.Date()
	y2, m2, d2 := day2.Date()

	hh1, mm1, ss1 := day1.Clock()
	hh2, mm2, ss2 := day2.Clock()

	day1 = time.Date(y1, m1, d1, hh1, mm1, ss1, 00, TIME_LOCATION)
	day2 = time.Date(y2, m2, d2, hh2, mm2, ss2, 00, TIME_LOCATION)

	return day1.Equal(day2)
}
