package controller

// TaskController stores methods for serving task requests
type TaskController struct{}

// // NewTaskAction adds a new task to database
// func (controller TaskController) NewTaskAction(w http.ResponseWriter, r *http.Request) {
// 	calendarID, _ := strconv.Atoi(r.FormValue("calendarId"))
// 	date := 0
// 	task := model.NewTaskModel(
// 		calendarID,
// 		r.FormValue("title"),
// 		date,
// 		r.FormValue("from"),
// 		r.FormValue("to"),
// 		r.FormValue("description"),
// 		false,
// 	)
// 	model.DefaultTaskRepository.AddTask(task)
// 	http.Redirect(w, r, "/", 303)
// }
