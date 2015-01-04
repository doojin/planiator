package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getDay_ShouldReturnCorrectDayIfItIsSmallerThanTen(t *testing.T) {
	day := getDay(1)

	assert.Equal(t, "01", day)
}

func Test_getDay_ShouldReturnCorrectDayIfItIsGreaterThanTen(t *testing.T) {
	day := getDay(15)

	assert.Equal(t, "15", day)
}

func Test_getMonth_ShouldReturnCorrectMonthIfItIsSmallerThanTen(t *testing.T) {
	date := time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)

	month := getMonth(date)

	assert.Equal(t, "01", month)
}

func Test_getMonth_ShouldReturnCorrectMonthIfItIsGreaterThanTen(t *testing.T) {
	date := time.Date(2001, time.December, 1, 0, 0, 0, 0, time.UTC)

	month := getMonth(date)

	assert.Equal(t, "12", month)
}

func Test_getYear_ShouldReturnCorrectYear(t *testing.T) {
	date := time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)

	year := getYear(date)

	assert.Equal(t, "2001", year)
}

func Test_GetDayId_ShouldReturnCorrectDayID(t *testing.T) {
	service := NewDateService()

	dayID := service.GetDayID("01", "01", "2001")

	assert.Equal(t, "01012001", dayID)
}

func Test_GetMonthDays_ShouldReturnCorrectMonthDaysWithNegativeOffset(t *testing.T) {
	service := NewDateService()

	days := service.GetMonthDays(0)

	assert.Equal(t, nil, days)
}
