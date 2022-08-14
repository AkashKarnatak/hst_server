package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mentor struct {
  Id primitive.ObjectID `json:"id" bson:"_id"`
  Name string `json:"mentor" bson:"Mentors"`
  Description string `json:"description" bson:"Description"`
  LinkedinProfile string `json:"linkedinProfile" bson:"Linked In profile"`
  Organization string `json:"organization" bson:"Organization"`
}
