package model

import (
	"math/rand"
	"time"
)

// TaskModel stores information about task
type TaskModel struct {
	ID          int    `bson:"uid"`
	CalendarID  int    `bson:"calendarId"`
	Title       string `bson:"title"`
	Date        int    `bson:"date"`
	From        string `bson:"from"`
	To          string `bson:"to"`
	Description string `bson:"desc"`
	IsDone      bool   `bson:"isDone"`
}

// NewTaskModel creates new instance of TaskModel
func NewTaskModel(calendarID int, title string, date int, from string, to string, description string, isDone bool) (model TaskModel) {
	model.CalendarID = calendarID
	model.Title = title
	model.Date = date
	model.From = from
	model.To = to
	model.Description = description
	model.IsDone = isDone

	rand.Seed(time.Now().Unix())
	model.ID = rand.Intn(9999999999-1000000000) + 1000000000

	return
}
