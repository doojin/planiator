package controller

import (
	"planiator/server/model"
	"planiator/server/service"
)

func populateCalendarList(data map[string]interface{}, userID int) {
	calendars := model.DefaultCalendarRepository.GetAll(userID)
	data["Calendars"] = calendars
}

func populateWeeks(data map[string]interface{}, monthOffset int) {
	weeks := service.DefaultDateService.GetWeeks(monthOffset)
	data["Weeks"] = weeks
}
