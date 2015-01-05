package service

import (
	"math"
	"strconv"
	"time"
)

// DateService contains methods for date manipulation
type DateService struct {
	TimeProvider TimeProviderI
}

// DefaultDateService is a default instance of DateService
var DefaultDateService = NewDateService()

// NewDateService returns a new instance of DateService
func NewDateService() (service DateService) {
	service.TimeProvider = timeProvider{}
	return
}

// GetMonthDays returns days of months
func (service DateService) GetMonthDays(monthOffset int) []string {
	days := []string{}
	now := service.TimeProvider.now()
	date := now.AddDate(0, monthOffset, 0)
	totalDays := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	for i := 1; i <= totalDays; i++ {
		days = append(days, service.GetDayID(service.getDay(i), service.getMonth(date), service.getYear(date)))
	}
	return days
}

// GetWeeks returns weeks of month
func (service DateService) GetWeeks(monthOffset int) [][7]map[string]interface{} {
	weeks := [][7]map[string]interface{}{}
	days := service.GetMonthDays(monthOffset)
	date := service.TimeProvider.now().AddDate(0, monthOffset, 0)
	firstDay := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	weekDay := service.dayOfWeek(firstDay)
	totalDays := len(days) + weekDay
	totalWeeks := int(math.Ceil(float64(totalDays) / 7.0))
	current := 0
	firstWeek := [7]map[string]interface{}{}
	for i := weekDay; i <= 6; i++ {
		firstWeek[i] = map[string]interface{}{"Id": days[current], "Day": current + 1}
		current++
	}
	weeks = append(weeks, firstWeek)
	for i := 1; i < totalWeeks; i++ {
		week := [7]map[string]interface{}{}
		for j := 0; j < 7; j++ {
			if current < len(days) {
				week[j] = map[string]interface{}{"Id": days[current], "Day": current + 1}
				current++
			}
		}
		weeks = append(weeks, week)
	}
	return weeks
}

// MonthName returns month name by it's offset
func (service DateService) MonthName(offset int) string {
	now := service.TimeProvider.now()
	then := now.AddDate(0, offset, 0)
	return then.Format("January")
}

// YearName returns year by month offset
func (service DateService) YearName(offset int) string {
	now := service.TimeProvider.now()
	then := now.AddDate(0, offset, 0)
	return then.Format("2006")
}

func (service DateService) getDay(day int) string {
	result := strconv.Itoa(day)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

func (service DateService) getMonth(date time.Time) string {
	month := strconv.Itoa(int(date.Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	return month
}

func (service DateService) dayOfWeek(day time.Time) int {
	weekDay := int(day.Weekday())
	weekDay--
	if weekDay == -1 {
		weekDay = 6
	}
	return weekDay
}

func (service DateService) getYear(date time.Time) string {
	return strconv.Itoa(date.Year())
}

// GetDayID returns day ID
func (service DateService) GetDayID(day string, month string, year string) string {
	return day + month + year
}

// TimeProviderI contains method signatures to work with time
type TimeProviderI interface {
	now() time.Time
}

type timeProvider struct {
}

func (provider timeProvider) now() time.Time {
	return time.Now()
}
