package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Startup struct {
  Id primitive.ObjectID `json:"id" bson:"_id"`
  Name string `json:"startupName" bson:"Startup Name"`
  Founder string `json:"founder" bson:"Name of Founder /cofounders"`
  EmailId string `json:"emailId" bson:"EmailId"`
  PhNo string `json:"phNo" bson:"phNo"`
  Description string `json:"description" bson:"Brief about Idea"`
  LinkedinProfile string `json:"linkedinProfile" bson:"Linkedin Profile"`
  Spoc string `json:"spoc" bson:"SPOC"`
}
