package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Meeting struct {
  Id primitive.ObjectID `json:"id" bson:"_id"`
  Name string `json:"name" bson:"Name"`
  Time string `json:"time" bson:"Time"`
  Day int `json:"day" bson:"Day"`
  Date string `json:"date" bson:"Date"`
  Venue string `json:"venue" bson:"Venue"`
  Description string `json:"description" bson:"Description"`
}
