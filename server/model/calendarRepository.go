package model

import (
	"planiator/server/conf"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DefaultCalendarRepository is a default instance of CalendarRepository
var DefaultCalendarRepository = NewCalendarRepository()

// CalendarRepositoryI contains method signatures to work with CalendarModel
type CalendarRepositoryI interface {
	AddCalendar(model CalendarModel)
}

// CalendarRepository containst methods to work with CalendarModel
type CalendarRepository struct {
	c *mgo.Collection
}

// NewCalendarRepository returns an exemplar of CalendarRepository
func NewCalendarRepository() (repo CalendarRepository) {
	repo.c = MongoSession.DB(conf.MongoDatabase).C("calendars")
	return
}

// AddCalendar adds new calendar to database
func (repo CalendarRepository) AddCalendar(model CalendarModel) {
	repo.c.Insert(model)
}

// GetAll returns all calendars from database
func (repo CalendarRepository) GetAll(userID int) (list []CalendarModel) {
	repo.c.Find(bson.M{"userId": userID}).All(&list)
	return
}

// UpdateCalendars removes all user calendars and replace them with a new list
func (repo CalendarRepository) UpdateCalendars(list []CalendarModel, userID int) {
	repo.c.RemoveAll(bson.M{"userId": userID})
	for _, calendar := range list {
		calendar.UserID = userID
		repo.c.Insert(calendar)
	}
}
