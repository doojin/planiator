package model

import (
	"planiator/server/conf"

	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
)

// MongoSession is a session for working with MongoDB
var MongoSession mgo.Session

func init() {
	logger := logging.MustGetLogger("MongoDB Session")
	MongoSession, err := mgo.Dial(conf.MongoURL)
	if err != nil {
		logger.Panic(err)
	}
	logger.Info("Connected to MongoDB server")
	defer MongoSession.Close()
}
