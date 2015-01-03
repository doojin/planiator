package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// getHomepageActionHandler

func Test_getHomepageActionHandler_ShouldReturnCorrectHandlerForSignInAction(t *testing.T) {
	action := "sign in"

	result := getHomepageActionHandler(action)

	assert.Equal(t, handleSignInAction, result)
}

func Test_getHomepageActionHandler_ShouldReturnCorrectHandlerForSignUpAction(t *testing.T) {
	action := "sign up"

	result := getHomepageActionHandler(action)

	assert.Equal(t, handleSignUpAction, result)
}

func Test_getHomepageActionHandler_ShouldReturnRedirectHandlerForUnknownAction(t *testing.T) {
	action := "dummy action name"

	result := getHomepageActionHandler(action)

	assert.Equal(t, handleHomepageRedirectAction, result)
}
