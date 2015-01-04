package form

import "planiator/server/model"

// SignInForm validates values of sign in form
type SignInForm struct {
	Email    string
	Password string

	userRepo model.UserRepositoryI
}

// NewSignInForm returns new NewSignInForm exemplar
func NewSignInForm(email string, password string) (form SignInForm) {
	form.Email = email
	form.Password = password

	form.userRepo = model.DefaultUserRepository
	return
}

// Validate validates sign up form
func (form SignInForm) Validate() bool {
	return form.userRepo.UserWithEmailAndPasswordExists(form.Email, form.Password)
}
