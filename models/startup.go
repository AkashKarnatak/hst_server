package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Startup struct {
  Id primitive.ObjectID `json:"id" bson:"_id"`
  Name string `json:"Startup Name" bson:"Startup Name"`
  Founder string `json:"Name of Founder /cofounders" bson:"Name of Founder /cofounders"`
  EmailId string `json:"EmailId" bson:"EmailId"`
  Description string `json:"Brief about Idea" bson:"Brief about Idea"`
  LinkedinProfile string `json:"Linkedin Profile" bson:"Linkedin Profile"`
  Spoc string `json:"SPOC" bson:"SPOC"`
}
