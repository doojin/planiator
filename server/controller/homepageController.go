package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"planiator/server/form"
	"planiator/server/messages"
	"planiator/server/model"
	"planiator/server/service"
	"planiator/server/session"

	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("Homepage Controller")

// HomepageController stores methods for serving homepage requests
type HomepageController struct {
}

// GetHomepage serves get request for homepage
func (hpc HomepageController) GetHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(service.DefaultAuthService.IsLoggedIn(r, w))
	fmt.Println(session.Get(r).Values)
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
	email, password := r.FormValue("email"), r.FormValue("password")
	signInForm := form.NewSignInForm(email, password)

	if ok := signInForm.Validate(); ok {
		user := model.DefaultUserRepository.FindByEmail(signInForm.Email)
		service.DefaultAuthService.Login(user.ID, r, w)
		http.Redirect(w, r, "/", 303)
	} else {
		tpl, err := template.ParseFiles("tpl/layout/homepage.tpl", "tpl/homepage.tpl")
		if err != nil {
			logger.Warning("Error while parsing template: %s", err)
		}
		tpl.Execute(w, map[string]interface{}{
			"AuthFailed":     messages.Messages["sign in form"]["ErrSignInWrongEmailOrPassword"],
			"SignInEmail":    signInForm.Email,
			"SignInPassword": signInForm.Password,
		})
	}
}

// handleSignUpAction is user registration handler
func handleSignUpAction(w http.ResponseWriter, r *http.Request) {
	email, password, passwordAgain := r.FormValue("email"), r.FormValue("password"), r.FormValue("password-again")
	signUpForm := form.NewSignUpForm(email, password, passwordAgain)

	if ok, _ := signUpForm.Validate(); !ok {
		data := map[string]interface{}{}
		data["EmailErrors"] = signUpForm.GetFieldErrMessages(form.SignUpFormEmailID)
		data["PasswordErrors"] = signUpForm.GetFieldErrMessages(form.SignUpFormPasswordID)
		data["PasswordAgainErrors"] = signUpForm.GetFieldErrMessages(form.SignUpFormPasswordAgainID)
		data["SignUpEmail"] = signUpForm.Email
		data["SignUpPassword"] = signUpForm.Password
		data["SignUpPasswordAgain"] = signUpForm.PasswordAgain

		tpl, err := template.ParseFiles("tpl/layout/homepage.tpl", "tpl/homepage.tpl")
		if err != nil {
			logger.Warning("Error while parsing template: %s", err)
		}
		tpl.Execute(w, data)
	} else {
		newUser := model.NewUserModel(signUpForm.Email, signUpForm.Password)
		model.DefaultUserRepository.AddUser(newUser)
		service.DefaultAuthService.Login(newUser.ID, r, w)
		http.Redirect(w, r, "/", 303)
	}
}

// handleHomepageRedirectAction redirects user to homepage
func handleHomepageRedirectAction(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 303)
}
