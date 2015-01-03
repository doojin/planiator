package service

import (
	"net/http"
	"planiator/server/session"
)

// AuthService manages log in / log out procedures
type AuthService struct {
}

var authService = new(AuthService)

// GetDefaultAuthService returns default AuthService
func GetDefaultAuthService() *AuthService {
	return authService
}

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
