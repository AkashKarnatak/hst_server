package models

type Meeting struct {
  MentorName string `json:"mentorName" bson:"mentorName"`
  MentorDescription string `json:"mentorDescription" bson:"mentorDescription"`
  StartupName string `json:"startupName" bson:"startupName"`
  StartupDescription string `json:"startupDescription" bson:"startupDescription"`
  Time string `json:"time" bson:"Time"`
  Day int `json:"day" bson:"Day"`
  Venue string `json:"venue" bson:"Venue"`
  Type string `json:"type" bson:"Type"`
  Panel string `json:"panel" bson:"Panel"`
}
