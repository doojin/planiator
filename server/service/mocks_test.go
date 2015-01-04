package service

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type timeProviderMock struct {
	mock.Mock
}

func (m *timeProviderMock) now() time.Time {
	args := m.Mock.Called()
	return args.Get(0).(time.Time)
}
