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

func main() {
	defer model.MongoSession.Close()
	userRepo := model.NewUserRepository()
	userRepo.EmailExists("aaa@aaa.aa")
	logger := logging.MustGetLogger("Web Server")
	r := mux.NewRouter()

	r.HandleFunc("/", homepageController.GetHomepage).Methods("GET")
	r.HandleFunc("/", homepageController.PostHomepage).Methods("POST")

	r.HandleFunc("/calendar", calendarController.GetCalendarPage).Methods("GET")

	r.HandleFunc("/logoff", logoffController.LogoffAction).Methods("GET")

	// Serving static files
	fileHandler := http.StripPrefix("/assets/", http.FileServer(http.Dir("../assets")))
	http.Handle("/assets/", fileHandler)

	logger.Info("Starting web server on port: %s", conf.Port)

	http.Handle("/", r)
	http.ListenAndServe(conf.ConnURL, nil)
}
