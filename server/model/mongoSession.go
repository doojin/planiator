package model

import (
	"planiator/server/conf"

	"gopkg.in/mgo.v2"
)

// MongoSession is a session for working with MongoDB
var MongoSession, _ = mgo.Dial(conf.MongoURL)
