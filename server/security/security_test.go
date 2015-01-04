package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncryptPassword_ShouldEncryptPasswordCorrectly(t *testing.T) {
	password := "dummy password"

	result := EncryptPassword(password)

	assert.Equal(t, "753487d47022011d5f6a2bd32d9b3bae", result)
}
