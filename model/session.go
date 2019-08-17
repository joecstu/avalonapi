package model

import "github.com/globalsign/mgo/bson"

type Session struct {
	Id    bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Email string
}
