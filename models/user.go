package models

type User struct {
  Id string `json:"id" bson:"_id"`
  EmailId string `json:"emailId" bson:"Email"`
  PhNo string `json:"phNo" bson:"Phone"`
  Tokens []string `json:"tokens" bson:"tokens"`
}

type Mentor struct {
  User `json:",inline" bson:",inline"`
  Name string `json:"name" bson:"Name"`
  Description string `json:"description" bson:"Description"`
  LinkedinProfile string `json:"linkedinProfile" bson:"LinkedIn"`
  Organization string `json:"organization" bson:"Organization"`
}

type Startup struct {
  User `json:",inline" bson:",inline"`
  Name string `json:"startupName" bson:"Name"`
  Founder string `json:"founder" bson:"Founder Name"`
  Description string `json:"description" bson:"Brief about Idea"`
  LinkedinProfile string `json:"linkedinProfile" bson:"Linkedin Profile"`
  Spoc string `json:"spoc" bson:"SPOC"`
}
