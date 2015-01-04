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
func (service DateService) GetMonthDays(monthOffset int) map[int]string {
	days := map[int]string{}
	now := service.TimeProvider.now()
	date := now.AddDate(0, monthOffset, 0)
	totalDays := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	for i := 1; i <= totalDays; i++ {
		days[i] = service.GetDayID(i, int(date.Month()), date.Year())
	}
	return days
}

// GetDayID returns day ID
func (service DateService) GetDayID(day int, month int, year int) int {
	dayIDString := strconv.Itoa(day) + strconv.Itoa(month) + strconv.Itoa(year)
	dayID, _ := strconv.Atoi(dayIDString)
	return dayID
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
