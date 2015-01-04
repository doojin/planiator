package controller

import (
	"net/http"
	"planiator/server/service"
)

// LogoffController stores method for serving log off request
type LogoffController struct{}

// LogoffAction loggs off user and redirects him to homepage
func (controller LogoffController) LogoffAction(w http.ResponseWriter, r *http.Request) {
	service.DefaultAuthService.LogOff(r, w)
	http.Redirect(w, r, "/", 303)
}
