package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int64         `json:"age" bson:"age"`
}


