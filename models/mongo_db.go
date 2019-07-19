package models

import "gopkg.in/mgo.v2/bson"

//DBFriends ... Database Model Friends
type DBFriends struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	UserID  string        `bson:"user_id,omitempty"`
	Friends []string      `bson:"friends,omitempty"`
}
