package form

import (
	"planiator/server/model"
	"regexp"
)

// SignUpForm validates values of sign up form
type SignUpForm struct {
	Email         string
	Password      string
	PasswordAgain string

	userRepo model.UserRepositoryI
}

// Field ids
const (
	EmailID         = 1
	PasswordID      = 2
	PasswordAgainID = 3
)

// Error codes
const (
	ErrSignUpEmailEmpty   = 100
	ErrSignUpEmailRegexp  = 101
	ErrSignUpEmailTooLong = 102
	ErrSignUpEmailExists  = 103

	ErrSignUpPasswordEmpty  = 200
	ErrSignUpPasswordLength = 201

	ErrSignUpPasswordAgainEmpty = 300
	ErrSignUpPasswordsDontMatch = 301
)

// Email
const (
	emailPattern   = "^\\S+@\\S+\\.\\S+$"
	emailMaxLength = 50
)

// Password
const (
	passwordMinLength = 6
	passwordMaxLength = 30
)

// NewSignUpForm returns new SignUpForm exemplar
func NewSignUpForm(email string, password string, passwordAgain string) (form SignUpForm) {
	form.Email = email
	form.Password = password
	form.PasswordAgain = passwordAgain

	form.userRepo = model.UserRepository{}
	return
}

// Validate valdates form data
func (form SignUpForm) Validate() (result bool, errCodes map[int][]int) {
	errCodes = map[int][]int{1: []int{}, 2: []int{}, 3: []int{}}
	result = true

	// Checking Email address
	if form.isEmailEmpty() {
		result = false
		addEmailError(errCodes, ErrSignUpEmailEmpty)

	} else {
		if !form.isEmailRegexpCorrect() {
			result = false
			addEmailError(errCodes, ErrSignUpEmailRegexp)
		}

		if !form.isEmailLengthCorrect() {
			result = false
			addEmailError(errCodes, ErrSignUpEmailTooLong)
		}

		if form.userRepo.EmailExists(form.Email) {
			result = false
			addEmailError(errCodes, ErrSignUpEmailExists)
		}
	}

	// Checking Password
	if form.isPasswordEmpty() {
		result = false
		addPasswordError(errCodes, ErrSignUpPasswordEmpty)
	} else {
		if !form.isPasswordLengthCorrect() {
			result = false
			addPasswordError(errCodes, ErrSignUpPasswordLength)
		}
	}

	// Checking "Password Again" field
	if form.isPasswordAgainEmpty() {
		result = false
		addPasswordAgainError(errCodes, ErrSignUpPasswordAgainEmpty)
	} else {
		if !form.passwordsMatch() {
			result = false
			addPasswordAgainError(errCodes, ErrSignUpPasswordsDontMatch)
		}
	}

	return result, errCodes
}

func (form SignUpForm) isEmailEmpty() bool {
	return len(form.Email) == 0
}

func (form SignUpForm) isEmailRegexpCorrect() bool {
	re := regexp.MustCompile(emailPattern)
	return re.MatchString(form.Email)
}

func (form SignUpForm) isEmailLengthCorrect() bool {
	length := len(form.Email)
	return length <= emailMaxLength
}

func (form SignUpForm) isPasswordEmpty() bool {
	return len(form.Password) == 0
}

func (form SignUpForm) isPasswordLengthCorrect() bool {
	length := len(form.Password)
	return length >= passwordMinLength && length <= passwordMaxLength
}

func (form SignUpForm) isPasswordAgainEmpty() bool {
	return len(form.PasswordAgain) == 0
}

func (form SignUpForm) passwordsMatch() bool {
	return form.Password == form.PasswordAgain
}

func addEmailError(errCodes map[int][]int, code int) {
	errCodes[EmailID] = append(errCodes[EmailID], code)
}

func addPasswordError(errCodes map[int][]int, code int) {
	errCodes[PasswordID] = append(errCodes[PasswordID], code)
}

func addPasswordAgainError(errCodes map[int][]int, code int) {
	errCodes[PasswordAgainID] = append(errCodes[PasswordAgainID], code)
}
