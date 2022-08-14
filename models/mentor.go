package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mentor struct {
  Id primitive.ObjectID `json:"id" bson:"_id"`
  Name string `json:"Mentors" bson:"Mentors"`
  Description string `json:"Description" bson:"Description"`
  LinkedinProfile string `json:"Linked In profile" bson:"Linked In profile"`
  Organization string `json:"Organization" bson:"Organization"`
}
