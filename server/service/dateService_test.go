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

func Test_GetMonthDays_ShouldReturnCorrectMonthDaysWithZeroOffset(t *testing.T) {
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC))
	service := NewDateService()
	service.TimeProvider = timeProvider

	days := service.GetMonthDays(0)

	expectedDays := []string{
		"01012001",
		"02012001",
		"03012001",
		"04012001",
		"05012001",
		"06012001",
		"07012001",
		"08012001",
		"09012001",
		"10012001",
		"11012001",
		"12012001",
		"13012001",
		"14012001",
		"15012001",
		"16012001",
		"17012001",
		"18012001",
		"19012001",
		"20012001",
		"21012001",
		"22012001",
		"23012001",
		"24012001",
		"25012001",
		"26012001",
		"27012001",
		"28012001",
		"29012001",
		"30012001",
		"31012001",
	}
	assert.Equal(t, expectedDays, days)
}

func Test_GetMonthDays_ShouldReturnCorrectMonthDaysWithNegativeOffset(t *testing.T) {
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC))
	service := NewDateService()
	service.TimeProvider = timeProvider

	days := service.GetMonthDays(-1)

	expectedDays := []string{
		"01122000",
		"02122000",
		"03122000",
		"04122000",
		"05122000",
		"06122000",
		"07122000",
		"08122000",
		"09122000",
		"10122000",
		"11122000",
		"12122000",
		"13122000",
		"14122000",
		"15122000",
		"16122000",
		"17122000",
		"18122000",
		"19122000",
		"20122000",
		"21122000",
		"22122000",
		"23122000",
		"24122000",
		"25122000",
		"26122000",
		"27122000",
		"28122000",
		"29122000",
		"30122000",
		"31122000",
	}
	assert.Equal(t, expectedDays, days)
}

func Test_GetMonthDays_ShouldReturnCorrectMonthDaysWithPositiveOffset(t *testing.T) {
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC))
	service := NewDateService()
	service.TimeProvider = timeProvider

	days := service.GetMonthDays(1)

	expectedDays := []string{
		"01022001",
		"02022001",
		"03022001",
		"04022001",
		"05022001",
		"06022001",
		"07022001",
		"08022001",
		"09022001",
		"10022001",
		"11022001",
		"12022001",
		"13022001",
		"14022001",
		"15022001",
		"16022001",
		"17022001",
		"18022001",
		"19022001",
		"20022001",
		"21022001",
		"22022001",
		"23022001",
		"24022001",
		"25022001",
		"26022001",
		"27022001",
		"28022001",
	}
	assert.Equal(t, expectedDays, days)
}
