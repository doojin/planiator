package form

import (
	"planiator/server/model"

	"github.com/stretchr/testify/mock"
)

type userRepositoryMock struct {
	mock.Mock
}

func (m *userRepositoryMock) EmailExists(email string) bool {
	args := m.Mock.Called(email)
	return args.Bool(0)
}

func (m *userRepositoryMock) UserWithEmailAndPasswordExists(email string, password string) bool {
	args := m.Mock.Called(email, password)
	return args.Bool(0)
}

func (m *userRepositoryMock) FindByEmail(email string) model.UserModel {
	args := m.Mock.Called(email)
	return args.Get(0).(model.UserModel)
}
