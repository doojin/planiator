package controller

import (
	"html/template"
	"net/http"

	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("Homepage Controller")

// HomepageController stores methods for serving homepage requests
type HomepageController struct {
}

// GetHomepage serves get request for homepage
func (hpc HomepageController) GetHomepage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("tpl/layout/homepage.tpl", "tpl/homepage.tpl")
	if err != nil {
		logger.Warning("Error while parsing template: %s", err)
	}
	tpl.Execute(w, map[string]interface{}{})
}

// PostHomepage determines action type and serves post request
func (hpc HomepageController) PostHomepage(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	handle := getHomepageActionHandler(action)
	handle(w, r)
}

func getHomepageActionHandler(action string) func(w http.ResponseWriter, r *http.Request) {
	switch action {
	case "sign in":
		{
			return handleSignInAction
		}
	case "sign up":
		{
			return handleSignUpAction
		}
	default:
		{
			return handleHomepageRedirectAction
		}
	}
}

// handleSignInAction is user authorization handler
func handleSignInAction(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("tpl/layout/homepage.tpl", "tpl/homepage.tpl")
	if err != nil {
		logger.Warning("Error while parsing template: %s", err)
	}
	tpl.Execute(w, nil)
}

// handleSignUpAction is user registration handler
func handleSignUpAction(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("tpl/layout/homepage.tpl", "tpl/homepage.tpl")
	if err != nil {
		logger.Warning("Error while parsing template: %s", err)
	}
	tpl.Execute(w, nil)
}

// handleHomepageRedirectAction redirects user to homepage
func handleHomepageRedirectAction(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 303)
}
