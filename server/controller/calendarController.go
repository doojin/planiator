package controller

import (
	"net/http"
	"text/template"
)

// CalendarController stores methods for serving calendar requests
type CalendarController struct{}

// GetCalendarPage serves get request for calendar page
func (controller CalendarController) GetCalendarPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles(
		"tpl/layout/default.tpl",
		"tpl/calendar.tpl",
		"tpl/blocks/navigation.tpl",
		"tpl/blocks/popups/calendar.tpl",
		"tpl/blocks/scripts/calendar.tpl",
		"tpl/blocks/styles/calendar.tpl",
	)
	if err != nil {
		logger.Warning("Error while parsing template: %s", err)
	}
	tpl.Execute(w, map[string]interface{}{})
}
