package main

import (
	"net/http"

	"github.com/doojin/planiator/server/conf"
	"github.com/doojin/planiator/server/controller"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
)

var homepageController controller.HomepageController

func init() {
	homepageController = controller.HomepageController{}
}

func main() {
	logger := logging.MustGetLogger("Web Server")
	r := mux.NewRouter()

	r.HandleFunc("/", homepageController.GetHomepage).Methods("GET")
	r.HandleFunc("/", homepageController.PostHomepage).Methods("POST")

	// Serving static files
	fileHandler := http.StripPrefix("/assets/", http.FileServer(http.Dir("../assets")))
	http.Handle("/assets/", fileHandler)

	logger.Info("Starting web server on port: %s", conf.Port)

	http.Handle("/", r)
	http.ListenAndServe(conf.ConnURL, nil)
}
