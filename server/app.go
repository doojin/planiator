package main

import (
	"net/http"
	"planiator/server/conf"
	"planiator/server/controller"
	"planiator/server/model"

	"github.com/gorilla/mux"
	"github.com/op/go-logging"
)

var homepageController = controller.HomepageController{}
var calendarController = controller.CalendarController{}
var logoffController = controller.LogoffController{}
var taskController = controller.TaskController{}

func main() {
	defer model.MongoSession.Close()
	logger := logging.MustGetLogger("Web Server")
	r := mux.NewRouter()

	r.HandleFunc("/", homepageController.GetHomepage).Methods("GET")
	r.HandleFunc("/", homepageController.PostHomepage).Methods("POST")

	r.HandleFunc("/calendar", calendarController.GetCalendarPage).Methods("GET")
	r.HandleFunc("/calendar/m{offset}", calendarController.GetCalendarPage).Methods("GET")
	r.HandleFunc("/new-calendar", calendarController.NewCalendarAction).Methods("POST")
	r.HandleFunc("/save-calendars", calendarController.UpdateCalendarsAction).Methods("POST")

	r.HandleFunc("/save-task", taskController.NewTaskAction).Methods("POST")

	r.HandleFunc("/logoff", logoffController.LogoffAction).Methods("GET")

	// Serving static files
	fileHandler := http.StripPrefix("/assets/", http.FileServer(http.Dir("../assets")))
	http.Handle("/assets/", fileHandler)

	logger.Info("Starting web server on port: %s", conf.Port)

	http.Handle("/", r)
	http.ListenAndServe(conf.ConnURL, nil)
}
