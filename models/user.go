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
  CurrentScheme string `json:"currentScheme" bson:"Current Scheme"`
  Sector string `json:"sector" bson:"Sector"`
  SubSector string `json:"subSector" bson:"Sub sector"`
  City string `json:"city" bson:"City"`
  ProductType string `json:"productType" bson:"Product Type"`
  MarketType string `json:"marketType" bson:"Market type"`
  CurrentProductStage string `json:"currentProductStage" bson:"Current Product Stage"`
  HimalayanFocused string `json:"himalayanFocused" bson:"Himalayan focussed"`
  HimachalBased string `json:"himachalBased" bson:"Himachal based"`
  StudentStartup string `json:"studentStartup" bson:"Student startup"`
  IITMandiStartup string `json:"IITMandiStartup" bson:"IIT Mandi Startup"`
  TrlCurrent string `json:"trlCurrent" bson:"TRL Current"`
  IncubationMode string `json:"incubationMode" bson:"Mode of Incubation"`
  Employees string `json:"employees" bson:"Employees"`
  InvestmentValue string `json:"investmentValue" bson:"Investment Value (L)"`
  FundsSanctioned string `json:"fundsSanctioned" bson:"Funds sanctioned"`
}
