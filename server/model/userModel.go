package model

import (
	"math/rand"
	"planiator/server/security"
	"time"
)

// UserModel stores information about user
type UserModel struct {
	ID           int    `bson:"uid"`
	Email        string `bson:"email"`
	PasswordHash string `bson:"password"`
}

// NewUserModel creates new instance of UserModel
func NewUserModel(email string, password string) UserModel {
	model := UserModel{}
	model.Email = email
	model.PasswordHash = security.EncryptPassword(password)

	rand.Seed(time.Now().Unix())
	model.ID = rand.Intn(9999999999-1000000000) + 1000000000

	return model
}
