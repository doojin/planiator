package model

import (
	"planiator/server/conf"
	"planiator/server/security"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DefaultUserRepository is a default instance of UserRepository
var DefaultUserRepository = NewUserRepository()

// UserRepositoryI contains method signatures to work with UserModel
type UserRepositoryI interface {
	EmailExists(email string) bool
	UserWithEmailAndPasswordExists(email string, password string) bool
	FindByEmail(email string) UserModel
}

// UserRepository containst methods to work with UserModel
type UserRepository struct {
	c *mgo.Collection
}

// NewUserRepository returns an exemplar of UserRepository
func NewUserRepository() (repo UserRepository) {
	repo.c = MongoSession.DB(conf.MongoDatabase).C("users")
	return
}

// EmailExists returns true if user with passed email exists in database
func (repo UserRepository) EmailExists(email string) bool {
	emailLowercase := strings.ToLower(email)
	userCount, _ := repo.c.Find(bson.M{"email": emailLowercase}).Count()
	return userCount > 0
}

// UserWithEmailAndPasswordExists returns true if user with passed email and password
// exists in database. Otherwise returns false
func (repo UserRepository) UserWithEmailAndPasswordExists(email string, password string) bool {
	encryptedPassword := security.EncryptPassword(password)
	userCount, _ := repo.c.Find(bson.M{"email": email, "password": encryptedPassword}).Count()
	return userCount > 0
}

// AddUser adds new user to database
func (repo UserRepository) AddUser(model UserModel) {
	model.Email = strings.ToLower(model.Email)
	repo.c.Insert(model)
}

// FindByEmail finds user by it's email
func (repo UserRepository) FindByEmail(email string) (model UserModel) {
	emailLowercase := strings.ToLower(email)
	repo.c.Find(bson.M{"email": emailLowercase}).One(&model)
	return
}
