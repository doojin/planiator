package controller

import "planiator/server/model"

func populateCalendarList(data map[string]interface{}, userID int) {
	calendars := model.DefaultCalendarRepository.GetAll(userID)
	data["Calendars"] = calendars
}

func populateDays(data map[string]interface{}, monthOffset int) {

}
