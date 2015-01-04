package model

import (
	"math/rand"
	"time"
)

// CalendarModel stores information about calendar
type CalendarModel struct {
	ID     int    `bson:"uid"`
	Name   string `bson:"name"`
	Color  string `bson:"color"`
	UserID int    `bson:"userId"`
}

// NewCalendarModel creates new instance of CalendarModel
func NewCalendarModel(name string, color string, userID int) (model CalendarModel) {
	model.Name = name
	model.Color = color
	model.UserID = userID

	rand.Seed(time.Now().Unix())
	model.ID = rand.Intn(9999999999-1000000000) + 1000000000

	return
}
