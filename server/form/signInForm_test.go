package form

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validate_ShouldReturnTrueIfUserWithEmailAndPasswordExists(t *testing.T) {
	userRepoMock := new(userRepositoryMock)
	userRepoMock.On("UserWithEmailAndPasswordExists", "dummyEmail", "dummyPassword").Return(true)
	signInForm := SignInForm{
		Email:    "dummyEmail",
		Password: "dummyPassword",
		userRepo: userRepoMock,
	}

	result := signInForm.Validate()

	assert.Equal(t, true, result)
}

func Test_Validate_ShouldReturnFalseIfUserWithEmailAndPasswordDoesntExistt(t *testing.T) {
	userRepoMock := new(userRepositoryMock)
	userRepoMock.On("UserWithEmailAndPasswordExists", "dummyEmail", "dummyPassword").Return(false)
	signInForm := SignInForm{
		Email:    "dummyEmail",
		Password: "dummyPassword",
		userRepo: userRepoMock,
	}

	result := signInForm.Validate()

	assert.Equal(t, false, result)
}
