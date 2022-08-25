package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AkashKarnatak/hst_server/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MeetingController struct { mappingColl *mongo.Collection }

func NewMeetingController(mappingColl *mongo.Collection) *MeetingController {
  return &MeetingController{mappingColl}
}

func (mc *MeetingController) GetMeetings(w http.ResponseWriter,
  r *http.Request, p httprouter.Params) {
  // identify whether user is mentor or startup
  var user string
  if p.ByName("id")[:3] == "STA" {
    user = "Startup"
  } else {
    user = "Mentor"
  }
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  var meetings []models.Meeting
  // query to fetch all meetings for user
  cursor, err := mc.mappingColl.Aggregate(
    ctx,
    bson.A{
      bson.M{
        "$match": bson.M{
          user + " Id": p.ByName("id"),
        },
      },
      bson.M{
        "$lookup": bson.M{
          "from": "mentors",
          "localField": "Mentor Id",
          "foreignField": "_id",
          "as": "mentorData",
        },
      },
      bson.M{
        "$unwind": bson.M{
          "path": "$mentorData",
          "preserveNullAndEmptyArrays": false,
        },
      },
      bson.M{
        "$lookup": bson.M{
          "from": "startups",
          "localField": "Startup Id",
          "foreignField": "_id",
          "as": "startupData",
        },
      },
      bson.M{
        "$unwind": bson.M{
          "path": "$startupData",
          "preserveNullAndEmptyArrays": false,
        },
      },
      bson.M{
        "$project": bson.M{
          "startupName": "$startupData.Name",
          "startupDescription": "$startupData.Descriptionn",
          "mentorName": "$mentorData.Name",
          "mentorDescription": "$mentorData.Descriptionn",
          "Time": 1,
          "Day": 1,
          "Venue": 1,
          "Type": 1,
          "Panel": 1,
        },
      },
    },
  )
  if err != nil {
    log.Printf("Error fetching events from db: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  if err = cursor.All(ctx, &meetings); err != nil {
    log.Printf("Error parsing events from db: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  // marshal and send meetings data as json
  mj, err := json.Marshal(meetings)
  if err != nil {
    log.Printf("Unable to marshal meetings to json: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n", mj)
}
