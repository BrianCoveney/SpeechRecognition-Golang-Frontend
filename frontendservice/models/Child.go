package models

import "gopkg.in/mgo.v2/bson"

type (
	Child struct {
		ID         bson.ObjectId  `bson:"_id" json:"id"`
		FirstName  string         `bson:"first_name" json:"first_name"`
		SecondName string         `bson:"second_name" json:"second_name"`
		Email      string         `bson:"email" json:"email"`
		Word       string         `bson:"word" json:"word"`
		Words      map[string]int `bson:"map_of_gliding_words" json:"map_of_gliding_words"`
	}
)
