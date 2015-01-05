package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_getDay_ShouldReturnCorrectDayIfItIsSmallerThanTen(t *testing.T) {
	day := DefaultDateService.getDay(1)

	assert.Equal(t, "01", day)
}

func Test_getDay_ShouldReturnCorrectDayIfItIsGreaterThanTen(t *testing.T) {
	day := DefaultDateService.getDay(15)

	assert.Equal(t, "15", day)
}

func Test_getMonth_ShouldReturnCorrectMonthIfItIsSmallerThanTen(t *testing.T) {
	date := time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)

	month := DefaultDateService.getMonth(date)

	assert.Equal(t, "01", month)
}

func Test_getMonth_ShouldReturnCorrectMonthIfItIsGreaterThanTen(t *testing.T) {
	date := time.Date(2001, time.December, 1, 0, 0, 0, 0, time.UTC)

	month := DefaultDateService.getMonth(date)

	assert.Equal(t, "12", month)
}

func Test_getYear_ShouldReturnCorrectYear(t *testing.T) {
	date := time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)

	year := DefaultDateService.getYear(date)

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

func Test_dayOfWeek_ShouldReturnCorrectDayOfWeekForMonday(t *testing.T) {
	date := time.Date(2015, time.January, 5, 0, 0, 0, 0, time.UTC)

	dayOfWeek := DefaultDateService.dayOfWeek(date)

	assert.Equal(t, 0, dayOfWeek)
}

func Test_dayOfWeek_ShouldReturnCorrectDayOfWeekForSunday(t *testing.T) {
	date := time.Date(2015, time.January, 4, 0, 0, 0, 0, time.UTC)

	dayOfWeek := DefaultDateService.dayOfWeek(date)

	assert.Equal(t, 6, dayOfWeek)
}

func Test_GetWeeks_ShouldReturnCorrectMonthWeeksWithZeroOffset(t *testing.T) {
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC))
	service := NewDateService()
	service.TimeProvider = timeProvider

	weeks := service.GetWeeks(0)

	assert.Equal(t, [][7]map[string]interface{}{
		[7]map[string]interface{}{
			3: map[string]interface{}{"Id": "01012015", "Day": 1},
			4: map[string]interface{}{"Id": "02012015", "Day": 2},
			5: map[string]interface{}{"Id": "03012015", "Day": 3},
			6: map[string]interface{}{"Id": "04012015", "Day": 4},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "05012015", "Day": 5},
			1: map[string]interface{}{"Id": "06012015", "Day": 6},
			2: map[string]interface{}{"Id": "07012015", "Day": 7},
			3: map[string]interface{}{"Id": "08012015", "Day": 8},
			4: map[string]interface{}{"Id": "09012015", "Day": 9},
			5: map[string]interface{}{"Id": "10012015", "Day": 10},
			6: map[string]interface{}{"Id": "11012015", "Day": 11},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "12012015", "Day": 12},
			1: map[string]interface{}{"Id": "13012015", "Day": 13},
			2: map[string]interface{}{"Id": "14012015", "Day": 14},
			3: map[string]interface{}{"Id": "15012015", "Day": 15},
			4: map[string]interface{}{"Id": "16012015", "Day": 16},
			5: map[string]interface{}{"Id": "17012015", "Day": 17},
			6: map[string]interface{}{"Id": "18012015", "Day": 18},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "19012015", "Day": 19},
			1: map[string]interface{}{"Id": "20012015", "Day": 20},
			2: map[string]interface{}{"Id": "21012015", "Day": 21},
			3: map[string]interface{}{"Id": "22012015", "Day": 22},
			4: map[string]interface{}{"Id": "23012015", "Day": 23},
			5: map[string]interface{}{"Id": "24012015", "Day": 24},
			6: map[string]interface{}{"Id": "25012015", "Day": 25},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "26012015", "Day": 26},
			1: map[string]interface{}{"Id": "27012015", "Day": 27},
			2: map[string]interface{}{"Id": "28012015", "Day": 28},
			3: map[string]interface{}{"Id": "29012015", "Day": 29},
			4: map[string]interface{}{"Id": "30012015", "Day": 30},
			5: map[string]interface{}{"Id": "31012015", "Day": 31},
		},
	}, weeks)
}

func Test_GetWeeks_ShouldReturnCorrectMonthWeeksWithNegativeOffset(t *testing.T) {
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC))
	service := NewDateService()
	service.TimeProvider = timeProvider

	weeks := service.GetWeeks(-1)

	assert.Equal(t, [][7]map[string]interface{}{
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "01122014", "Day": 1},
			1: map[string]interface{}{"Id": "02122014", "Day": 2},
			2: map[string]interface{}{"Id": "03122014", "Day": 3},
			3: map[string]interface{}{"Id": "04122014", "Day": 4},
			4: map[string]interface{}{"Id": "05122014", "Day": 5},
			5: map[string]interface{}{"Id": "06122014", "Day": 6},
			6: map[string]interface{}{"Id": "07122014", "Day": 7},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "08122014", "Day": 8},
			1: map[string]interface{}{"Id": "09122014", "Day": 9},
			2: map[string]interface{}{"Id": "10122014", "Day": 10},
			3: map[string]interface{}{"Id": "11122014", "Day": 11},
			4: map[string]interface{}{"Id": "12122014", "Day": 12},
			5: map[string]interface{}{"Id": "13122014", "Day": 13},
			6: map[string]interface{}{"Id": "14122014", "Day": 14},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "15122014", "Day": 15},
			1: map[string]interface{}{"Id": "16122014", "Day": 16},
			2: map[string]interface{}{"Id": "17122014", "Day": 17},
			3: map[string]interface{}{"Id": "18122014", "Day": 18},
			4: map[string]interface{}{"Id": "19122014", "Day": 19},
			5: map[string]interface{}{"Id": "20122014", "Day": 20},
			6: map[string]interface{}{"Id": "21122014", "Day": 21},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "22122014", "Day": 22},
			1: map[string]interface{}{"Id": "23122014", "Day": 23},
			2: map[string]interface{}{"Id": "24122014", "Day": 24},
			3: map[string]interface{}{"Id": "25122014", "Day": 25},
			4: map[string]interface{}{"Id": "26122014", "Day": 26},
			5: map[string]interface{}{"Id": "27122014", "Day": 27},
			6: map[string]interface{}{"Id": "28122014", "Day": 28},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "29122014", "Day": 29},
			1: map[string]interface{}{"Id": "30122014", "Day": 30},
			2: map[string]interface{}{"Id": "31122014", "Day": 31},
		},
	}, weeks)
}

func Test_GetWeeks_ShouldReturnCorrectMonthWeeksWithPositiveOffset(t *testing.T) {
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC))
	service := NewDateService()
	service.TimeProvider = timeProvider

	weeks := service.GetWeeks(1)

	assert.Equal(t, [][7]map[string]interface{}{
		[7]map[string]interface{}{
			6: map[string]interface{}{"Id": "01022015", "Day": 1},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "02022015", "Day": 2},
			1: map[string]interface{}{"Id": "03022015", "Day": 3},
			2: map[string]interface{}{"Id": "04022015", "Day": 4},
			3: map[string]interface{}{"Id": "05022015", "Day": 5},
			4: map[string]interface{}{"Id": "06022015", "Day": 6},
			5: map[string]interface{}{"Id": "07022015", "Day": 7},
			6: map[string]interface{}{"Id": "08022015", "Day": 8},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "09022015", "Day": 9},
			1: map[string]interface{}{"Id": "10022015", "Day": 10},
			2: map[string]interface{}{"Id": "11022015", "Day": 11},
			3: map[string]interface{}{"Id": "12022015", "Day": 12},
			4: map[string]interface{}{"Id": "13022015", "Day": 13},
			5: map[string]interface{}{"Id": "14022015", "Day": 14},
			6: map[string]interface{}{"Id": "15022015", "Day": 15},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "16022015", "Day": 16},
			1: map[string]interface{}{"Id": "17022015", "Day": 17},
			2: map[string]interface{}{"Id": "18022015", "Day": 18},
			3: map[string]interface{}{"Id": "19022015", "Day": 19},
			4: map[string]interface{}{"Id": "20022015", "Day": 20},
			5: map[string]interface{}{"Id": "21022015", "Day": 21},
			6: map[string]interface{}{"Id": "22022015", "Day": 22},
		},
		[7]map[string]interface{}{
			0: map[string]interface{}{"Id": "23022015", "Day": 23},
			1: map[string]interface{}{"Id": "24022015", "Day": 24},
			2: map[string]interface{}{"Id": "25022015", "Day": 25},
			3: map[string]interface{}{"Id": "26022015", "Day": 26},
			4: map[string]interface{}{"Id": "27022015", "Day": 27},
			5: map[string]interface{}{"Id": "28022015", "Day": 28},
		},
	}, weeks)
}

func Test_MonthName_ShouldReturnCorrectMonthNameWithZeroOffset(t *testing.T) {
	date := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(date)
	service := NewDateService()
	service.TimeProvider = timeProvider

	monthName := service.MonthName(0)

	assert.Equal(t, "January", monthName)
}

func Test_MonthName_ShouldReturnCorrectMonthNameWithNegativeOffset(t *testing.T) {
	date := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(date)
	service := NewDateService()
	service.TimeProvider = timeProvider

	monthName := service.MonthName(-1)

	assert.Equal(t, "December", monthName)
}

func Test_MonthName_ShouldReturnCorrectMonthNameWithPositiveOffset(t *testing.T) {
	date := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(date)
	service := NewDateService()
	service.TimeProvider = timeProvider

	monthName := service.MonthName(1)

	assert.Equal(t, "February", monthName)
}

func Test_YearName_ShouldReturnCorrectMonthNameWithZeroOffset(t *testing.T) {
	date := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(date)
	service := NewDateService()
	service.TimeProvider = timeProvider

	monthName := service.YearName(0)

	assert.Equal(t, "2015", monthName)
}

func Test_YearName_ShouldReturnCorrectMonthNameWithNegativeOffset(t *testing.T) {
	date := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(date)
	service := NewDateService()
	service.TimeProvider = timeProvider

	monthName := service.YearName(-1)

	assert.Equal(t, "2014", monthName)
}

func Test_YearName_ShouldReturnCorrectMonthNameWithPositiveOffset(t *testing.T) {
	date := time.Date(2015, time.December, 1, 0, 0, 0, 0, time.UTC)
	timeProvider := new(timeProviderMock)
	timeProvider.On("now").Return(date)
	service := NewDateService()
	service.TimeProvider = timeProvider

	monthName := service.YearName(1)

	assert.Equal(t, "2016", monthName)
}
