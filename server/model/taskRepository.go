package model

import (
	"planiator/server/conf"

	"gopkg.in/mgo.v2"
)

// DefaultTaskRepository is a default instance of TaskRepository
var DefaultTaskRepository = NewTaskRepository()

// DefaultTaskRepositoryI contains method signatures to work with TaskModel
type DefaultTaskRepositoryI interface {
}

// TaskRepository containst methods to work with TaskModel
type TaskRepository struct {
	c *mgo.Collection
}

// NewTaskRepository returns an exemplar of TaskRepository
func NewTaskRepository() (repository TaskRepository) {
	repository.c = MongoSession.DB(conf.MongoDatabase).C("tasks")
	return
}

// AddTask adds tasks to database
func (repository TaskRepository) AddTask(model TaskModel) {
	repository.c.Insert(model)
}
