package service

import (
	"net/http"
	"planiator/server/session"
)

// AuthService manages log in / log out procedures
type AuthService struct {
}

// DefaultService is a default instancce of AuthService
var DefaultAuthService = new(AuthService)

// Login authorizes user in system
func (service AuthService) Login(id int, r *http.Request, w http.ResponseWriter) {
	sess := session.Get(r)
	sess.Values["userId"] = id
	sess.Save(r, w)
}

// IsLoggedIn returns true if user is authorized and
// returns false if not
func (service AuthService) IsLoggedIn(r *http.Request, w http.ResponseWriter) bool {
	sess := session.Get(r)
	return sess.Values["userId"] != nil
}

// LogOff loggs user off
func (service AuthService) LogOff(r *http.Request, w http.ResponseWriter) {
	sess := session.Get(r)
	delete(sess.Values, "userId")
	sess.Save(r, w)
}

// GetUserID returns id of actual user
func (service AuthService) GetUserID(r *http.Request, w http.ResponseWriter) int {
	sess := session.Get(r)
	return sess.Values["userId"].(int)
}
