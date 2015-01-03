package session

import (
	"net/http"
	"planiator/server/conf"

	"github.com/gorilla/sessions"
)

var store = sessions.NewFilesystemStore("session/session_files", []byte(conf.CookieStoreKey))

// Get returns default session
func Get(r *http.Request) *sessions.Session {
	sess, _ := store.Get(r, "default")
	return sess
}
