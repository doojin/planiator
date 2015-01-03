package model

import (
	"planiator/server/conf"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserRepositoryI contains method signatures to work with UserModel
type UserRepositoryI interface {
	EmailExists(email string) bool
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
	userCount, _ := repo.c.Find(bson.M{"email": email}).Count()
	return userCount > 0
}
