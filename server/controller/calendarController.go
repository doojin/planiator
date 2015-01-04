package controller

import (
	"encoding/json"
	"html/template"
	"net/http"
	"planiator/server/model"
	"planiator/server/service"
)

// CalendarController stores methods for serving calendar requests
type CalendarController struct{}

// GetCalendarPage serves get request for calendar page
func (controller CalendarController) GetCalendarPage(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	userID := service.DefaultAuthService.GetUserID(r, w)
	populateCalendarList(data, userID)

	tpl, err := template.ParseFiles(
		"tpl/layout/default.tpl",
		"tpl/calendar.tpl",
		"tpl/blocks/navigation.tpl",
		"tpl/blocks/popups/calendar.tpl",
		"tpl/blocks/popups/myCalendarsPopup.tpl",
		"tpl/blocks/scripts/calendar.tpl",
		"tpl/blocks/styles/calendar.tpl",
	)
	if err != nil {
		logger.Warning("Error while parsing template: %s", err)
	}
	tpl.Execute(w, data)
}

// NewCalendarAction adds a new calendar to database
func (controller CalendarController) NewCalendarAction(w http.ResponseWriter, r *http.Request) {
	userID := service.DefaultAuthService.GetUserID(r, w)
	newCalendar := model.NewCalendarModel("Untitled Calendar", "#464D4B", userID)
	model.DefaultCalendarRepository.AddCalendar(newCalendar)
	jsonResponse, _ := json.Marshal(newCalendar)
	w.Write(jsonResponse)
}

// UpdateCalendarsAction overrides all user calendars
func (controller CalendarController) UpdateCalendarsAction(w http.ResponseWriter, r *http.Request) {
	list := []model.CalendarModel{}
	userID := service.DefaultAuthService.GetUserID(r, w)
	calendars := r.FormValue("calendars")
	json.Unmarshal([]byte(calendars), &list)
	model.DefaultCalendarRepository.UpdateCalendars(list, userID)
}
