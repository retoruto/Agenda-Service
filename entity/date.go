package entity

import (
	"fmt"
	"strconv"
)

type Date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
}

func (m *Date) initDate(t_year, t_month, t_day, t_hour, t_minute int) {
	m.Year = t_year
	m.Month = t_month
	m.Day = t_day
	m.Hour = t_hour
	m.Minute = t_minute
}

func (m Date) getYear() int {
	return m.Year
}
func (m *Date) setYear(y int) {
	m.Year = y
}
func (m Date) getMonth() int {
	return m.Month
}
func (m *Date) setMonth(mo int) {
	m.Month = mo
}
func (m Date) getDay() int {
	return m.Day
}
func (m *Date) setDay(d int) {
	m.Day = d
}
func (m Date) getHour() int {
	return m.Hour
}
func (m *Date) setHour(h int) {
	m.Hour = h
}
func (m Date) getMinute() int {
	return m.Minute
}
func (m *Date) setMinute(mi int) {
	m.Minute = mi
}
func (t_date Date) isValid() bool {
	if t_date.Year < 1000 || t_date.Year > 9999 {
		return false
	}
	if t_date.Month < 1 || t_date.Month > 12 {
		return false
	}
	if t_date.Day < 1 || t_date.Day > 31 {
		return false
	}
	if t_date.Hour < 0 || t_date.Hour > 23 {
		return false
	}
	if t_date.Minute < 0 || t_date.Minute > 59 {
		return false
	}
	if ((t_date.Month%2 == 1 && t_date.Month < 8) || (t_date.Month%2 == 0 && t_date.Month >= 8)) && (t_date.Day > 31 || t_date.Day < 1) {
		return false
	}
	if (t_date.Month == 4 || t_date.Month == 6 || t_date.Month == 9 || t_date.Month == 11) && (t_date.Day > 30 || t_date.Day < 1) {
		return false
	}
	if ((t_date.Year%4 == 0 && t_date.Year%100 != 0) || t_date.Year%400 == 0) && (t_date.Day > 29 || t_date.Day < 1) && t_date.Month == 2 {
		return false
	}
	if !((t_date.Year%4 == 0 && t_date.Year%100 != 0) || t_date.Year%400 == 0) && (t_date.Day > 28 || t_date.Day < 1) && t_date.Month == 2 {
		return false
	}
	return true
}

func stringToInt(s string) int {
	result, error := strconv.Atoi(s)

	if error != nil {
		fmt.Println("fail")
	}
	return result
}
//"0000-00-00/00:00"
func stringToDate(t_dateString string) Date {
	var i int
	var x Date
	for i = 0; i < 4; i++ {
		if t_dateString[i] > '9' || t_dateString[i] < '0' {
			x.initDate(0, 0, 0, 0, 0)
			return x
		}
	}
	if t_dateString[4] != '-' || t_dateString[7] != '-' || len(t_dateString) != 16 {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if t_dateString[10] != '/' || t_dateString[13] != ':' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if t_dateString[14] > '9' || t_dateString[14] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if t_dateString[15] > '9' || t_dateString[15] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if t_dateString[5] > '9' || t_dateString[5] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if t_dateString[6] > '9' || t_dateString[6] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if t_dateString[8] > '9' || t_dateString[8] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if t_dateString[9] > '9' || t_dateString[9] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if t_dateString[11] > '9' || t_dateString[11] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	if t_dateString[12] > '9' || t_dateString[12] < '0' {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}
	//0000-00-00/00:00
	x.setYear(stringToInt(t_dateString[0:4]))
	x.setMonth(stringToInt(t_dateString[5:7]))
	x.setDay(stringToInt(t_dateString[8:10]))
	x.setHour(stringToInt(t_dateString[11:13]))
	x.setMinute(stringToInt(t_dateString[14:16]))
	if x.isValid() != false {
		return x
	} else {
		x.initDate(0, 0, 0, 0, 0)
		return x
	}

}
func IntToString(a int) string {
	var result_string string
	result_string = strconv.Itoa(a)
	return result_string
}
func dateToString(t_date Date) string {
	if t_date.isValid() == false {
		return "0000-00-00/00:00"
	}

	var re string = ""
	re += IntToString(t_date.Year) + "-"

	if t_date.Month < 10 {
		re += "0"
	}
	re += IntToString(t_date.Month) + "-"

	if t_date.Day < 10 {
		re += "0"
	}
	re += IntToString(t_date.Day) + "/"

	if t_date.Hour < 10 {
		re += "0"
	}
	re += IntToString(t_date.Hour) + ":"

	if t_date.Minute < 10 {
		re += "0"
	}
	re += IntToString(t_date.Minute)
	return re
}

func (m *Date) copyDate(t Date) {
	m.Year = t.Year
	m.Month = t.Month
	m.Day = t.Day
	m.Hour = t.Hour
	m.Minute = t.Minute
}

func (m_date Date) isTheSame(t_date Date) bool {
	return (t_date.getYear() == m_date.getYear() &&
		t_date.getMonth() == m_date.getMonth() &&
		t_date.getDay() == m_date.getDay() &&
		t_date.getHour() == m_date.getHour() &&
		t_date.getMinute() == m_date.getMinute())
}

func (m Date) isMoreThan(t_date Date) bool {
	if m.Year < t_date.Year {
		return false
	}
	if m.Year > t_date.Year {
		return true
	}
	if m.Month < t_date.Month {
		return false
	}
	if m.Month > t_date.Month {
		return true
	}
	if m.Day < t_date.Day {
		return false
	}
	if m.Day > t_date.Day {
		return true
	}
	if m.Hour < t_date.Hour {
		return false
	}
	if m.Hour > t_date.Hour {
		return true
	}
	if m.Minute < t_date.Minute {
		return false
	}
	if m.Minute > t_date.Minute {
		return true
	}
	return false
}
func (m Date) isLessThan(t_date Date) bool {
	if m.isMoreThan(t_date) != true && m.isTheSame(t_date) != true {
		return true
	} else {
		return false
	}
}

func (m Date) GreaterOrEqual(t_date Date) bool {
	if m.isLessThan(t_date) != true {
		return true
	} else {
		return false
	}
}

func (m Date) SmallerOrEqual(t_date Date) bool {
	if m.isMoreThan(t_date) != true {
		return true
	} else {
		return false
	}
}
