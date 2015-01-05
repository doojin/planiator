package controller

import (
	"encoding/json"
	"net/http"
	"planiator/server/model"
	"strconv"
)

// TaskController stores methods for serving task requests
type TaskController struct{}

// NewTaskAction adds a new task to database
func (controller TaskController) NewTaskAction(w http.ResponseWriter, r *http.Request) {
	calendarID, _ := strconv.Atoi(r.FormValue("calendarId"))
	task := model.NewTaskModel(
		calendarID,
		r.FormValue("title"),
		r.FormValue("dayId"),
		r.FormValue("from"),
		r.FormValue("to"),
		r.FormValue("desc"),
		false,
	)
	model.DefaultTaskRepository.AddTask(task)
	jsonResponse, _ := json.Marshal(task)
	w.Write(jsonResponse)
}
