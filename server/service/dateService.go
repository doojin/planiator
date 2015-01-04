package service

import (
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
		days = append(days, service.GetDayID(getDay(i), getMonth(date), getYear(date)))
	}
	return days
}

func getDay(day int) string {
	result := strconv.Itoa(day)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

func getMonth(date time.Time) string {
	month := strconv.Itoa(int(date.Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	return month
}

func getYear(date time.Time) string {
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
