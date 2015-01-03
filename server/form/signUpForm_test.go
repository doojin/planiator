package form

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type userRepositoryMock struct {
	mock.Mock
}

func (m *userRepositoryMock) EmailExists(email string) bool {
	args := m.Mock.Called(email)
	return args.Bool(0)
}

// isEmailEmpty

func Test_isEmailEmpty_ShouldReturnTrueIfEmailIsEmpty(t *testing.T) {
	form := SignUpForm{Email: ""}

	result := form.isEmailEmpty()

	assert.Equal(t, true, result)
}

func Test_isEmailEmpty_ShouldReturnFalseIfEmailIsNotEmpty(t *testing.T) {
	form := SignUpForm{Email: "dummy email"}

	result := form.isEmailEmpty()

	assert.Equal(t, false, result)
}

// isEmailRegexpCorrect

func Test_isEmailRegexpCorrect_ShouldReturnTrueIfEmailRegexpIsCorrect(t *testing.T) {
	forms := []SignUpForm{}
	forms = append(forms, SignUpForm{Email: "aaa@bbb.ccc"})
	forms = append(forms, SignUpForm{Email: "a@b.c"})
	forms = append(forms, SignUpForm{Email: "123@456.789"})
	forms = append(forms, SignUpForm{Email: "1@2.3"})
	forms = append(forms, SignUpForm{Email: "a1@b2.c3"})
	forms = append(forms, SignUpForm{Email: "*@*.*"})
	forms = append(forms, SignUpForm{Email: "@@@.."})

	for _, form := range forms {
		assert.Equal(t, true, form.isEmailRegexpCorrect())
	}
}

func Test_isEmailRegexpCorrect_ShouldReturnFalseIfEmailRegexpIsIncorrect(t *testing.T) {
	forms := []SignUpForm{}
	forms = append(forms, SignUpForm{Email: "aaa.ccc"})
	forms = append(forms, SignUpForm{Email: "aaa@bbb"})
	forms = append(forms, SignUpForm{Email: "@bbb.ccc"})
	forms = append(forms, SignUpForm{Email: "aaa@.ccc"})

	for _, form := range forms {
		assert.Equal(t, false, form.isEmailRegexpCorrect())
	}
}

// isEmailLengthCorrect

func Test_isEmailLengthCorrect_ShouldReturnTrueIfEmailLengthIsCorrect(t *testing.T) {
	forms := []SignUpForm{}
	forms = append(forms, SignUpForm{Email: "11111111112222222222333333333344444444445555555555"})
	forms = append(forms, SignUpForm{Email: "1111111111222222222233333333334444444444555555555"})

	for _, form := range forms {
		assert.Equal(t, true, form.isEmailLengthCorrect())
	}
}

func Test_isEmailLengthCorrect_ShouldReturnFalseIfEmailLengthIsIncorrect(t *testing.T) {
	form := SignUpForm{Email: "111111111122222222223333333333444444444455555555556"}

	result := form.isEmailLengthCorrect()

	assert.Equal(t, false, result)
}

// isPasswordEmpty

func Test_isPasswordEmpty_ShouldReturnTrueIfPasswordIsEmpty(t *testing.T) {
	form := SignUpForm{Password: ""}

	result := form.isPasswordEmpty()

	assert.Equal(t, true, result)
}

func Test_isPasswordEmpty_ShouldReturnFalseIfPasswordIsNotEmpty(t *testing.T) {
	form := SignUpForm{Password: "dummy password"}

	result := form.isPasswordEmpty()

	assert.Equal(t, false, result)
}

// isPasswordLengthCorrect

func Test_isPasswordLengthCorrect_ShouldReturnTrueIfPasswordLengthIsCorrect(t *testing.T) {
	forms := []SignUpForm{}
	forms = append(forms, SignUpForm{Password: "111111111122222222223333333333"})
	forms = append(forms, SignUpForm{Password: "11111111112222222222333333333"})
	forms = append(forms, SignUpForm{Password: "1234567"})
	forms = append(forms, SignUpForm{Password: "123456"})

	for _, form := range forms {
		assert.Equal(t, true, form.isPasswordLengthCorrect())
	}
}

func Test_isPasswordLengthCorrect_ShouldReturnFalseIfPasswordLengthIsIncorrect(t *testing.T) {
	forms := []SignUpForm{}
	forms = append(forms, SignUpForm{Password: "1111111111222222222233333333334"})
	forms = append(forms, SignUpForm{Password: "12345"})

	for _, form := range forms {
		assert.Equal(t, false, form.isPasswordLengthCorrect())
	}
}

// isPasswordAgainEmpty

func Test_isPasswordAgainEmpty_ShouldReturnTrueIfPasswordAgainIsEmpty(t *testing.T) {
	form := SignUpForm{PasswordAgain: ""}

	result := form.isPasswordAgainEmpty()

	assert.Equal(t, true, result)
}

func Test_isPasswordAgainEmpty_ShouldReturnFalseIfPasswordAgainIsNotEmpty(t *testing.T) {
	form := SignUpForm{PasswordAgain: "dummy password again"}

	result := form.isPasswordAgainEmpty()

	assert.Equal(t, false, result)
}

// passwordsMatch

func Test_passwordsMatch_ShouldReturnTrueIfPasswordsMatch(t *testing.T) {
	forms := []SignUpForm{}
	forms = append(forms, SignUpForm{Password: "abcdefg", PasswordAgain: "abcdefg"})
	forms = append(forms, SignUpForm{Password: "", PasswordAgain: ""})

	for _, form := range forms {
		assert.Equal(t, true, form.passwordsMatch())
	}
}

func Test_passwordsMatch_ShouldReturnFalseIfPasswordsDontMatch(t *testing.T) {
	forms := []SignUpForm{}
	forms = append(forms, SignUpForm{Password: "abcdefg", PasswordAgain: "abcdefgh"})
	forms = append(forms, SignUpForm{Password: "", PasswordAgain: " "})

	for _, form := range forms {
		assert.Equal(t, false, form.passwordsMatch())
	}
}

// addEmailError

func Test_addEmailError_ShouldAddEmailErrorCode(t *testing.T) {
	errCodes := map[int][]int{}

	addEmailError(errCodes, 1)
	addEmailError(errCodes, 2)

	assert.Equal(t, len(errCodes), 1)
	assert.Equal(t, len(errCodes[EmailID]), 2)
	assert.Equal(t, errCodes, map[int][]int{0: []int{1, 2}})
}

// addPasswordError

func Test_addPasswordError_ShouldAddPasswordErrorCode(t *testing.T) {
	errCodes := map[int][]int{}

	addPasswordError(errCodes, 1)
	addPasswordError(errCodes, 2)

	assert.Equal(t, len(errCodes), 1)
	assert.Equal(t, len(errCodes[PasswordID]), 2)
	assert.Equal(t, errCodes, map[int][]int{1: []int{1, 2}})
}

// addPasswordAgainError

func Test_addPasswordAgainError_ShouldAddPasswordAgainErrorCode(t *testing.T) {
	errCodes := map[int][]int{}

	addPasswordAgainError(errCodes, 1)
	addPasswordAgainError(errCodes, 2)

	assert.Equal(t, len(errCodes), 1)
	assert.Equal(t, len(errCodes[PasswordAgainID]), 2)
	assert.Equal(t, errCodes, map[int][]int{2: []int{1, 2}})
}

// Validate

func Test_Validate_ShouldReturnOnlyEmptyFieldErrorCodesIfFieldsAreEmpty(t *testing.T) {
	form := SignUpForm{}

	result, errCodes := form.Validate()

	assert.Equal(t, false, result)
	assert.Equal(t, map[int][]int{0: []int{0}, 1: []int{4}, 2: []int{6}}, errCodes)
}

func Test_Validate_ShouldAddErrorsCorrectlyWhenFieldsAreNotEmpty(t *testing.T) {
	mockedUserRepo := new(userRepositoryMock)
	mockedUserRepo.On("EmailExists", "11111111112222222222333333333344444444445555555555@").Return(true)
	email := "11111111112222222222333333333344444444445555555555@"
	password := "11111"
	passwordAgain := "11112"
	form := SignUpForm{Email: email, Password: password, PasswordAgain: passwordAgain}
	form.userRepo = mockedUserRepo

	result, errCodes := form.Validate()

	assert.Equal(t, false, result)
	assert.Equal(t, map[int][]int{0: []int{1, 2, 3}, 1: []int{5}, 2: []int{7}}, errCodes)
}
