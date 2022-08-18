package controllers

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
//
// 	"github.com/AkashKarnatak/hst_server/middlewares"
// 	"github.com/AkashKarnatak/hst_server/models"
// 	"github.com/julienschmidt/httprouter"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )
//
// type MeetingController struct {
//   coll *mongo.Collection
// }
//
// func NewMeetingController(coll *mongo.Collection) *MeetingController {
//   return &MeetingController{coll}
// }
//
// func (mc *MeetingController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//   // let's find the meeting with matching phNo and email in the request
//   ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
//   defer cancel()
//   var meeting models.Meeting
//   err := mc.coll.FindOne(
//     ctx,
//     bson.M{ "EmailId": r.FormValue("EmailId"), "phNo": r.FormValue("phNo")},
//   ).Decode(&meeting)
//   if err != nil {
//     log.Fatalf("Unable to fetch and decode user information\n%v\n", err)
//     w.WriteHeader(http.StatusInternalServerError)
//     fmt.Fprintln(w, "Internal server error")
//     return
//   }
//   tokenString, err := middlewares.GenerateNewToken(meeting.Id)
//   if err != nil {
//     log.Fatalf("Unable to create new token\n%v\n", err)
//     w.WriteHeader(http.StatusInternalServerError)
//     fmt.Fprintln(w, "Internal server error")
//     return
//   }
//   meeting.Tokens = append(meeting.Tokens, tokenString)
//   // append token
//   // update db
// }
//
// func (mc *MeetingController) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//   // find user
//   // append token
//   // update db
// }
//
// func (mc *MeetingController) LogoutAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//   // find user
//   // append token
//   // update db
// }
//
// func (mc *MeetingController) GetMeetings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//   ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
//   defer cancel()
//   cursor, err := mc.coll.Find(ctx, bson.M{})
//   if err != nil {
//     log.Fatalf("Error in retrieving data\n%v\n", err)
//     w.WriteHeader(http.StatusInternalServerError)
//     fmt.Fprintln(w, "Internal server error")
//     return
//   }
//   var res []models.Meeting
//   err = cursor.All(ctx, &res)
//   if err != nil {
//     log.Fatalf("Unable to parse collection data\n%v\n", err)
//     w.WriteHeader(http.StatusInternalServerError)
//     fmt.Fprintln(w, "Internal server error")
//     return
//   }
//   resJson, err := json.Marshal(res)
//   if err != nil {
//     log.Fatalf("Unable to marshal data to json\n%v\n", err)
//     w.WriteHeader(http.StatusInternalServerError)
//     fmt.Fprintln(w, "Internal server error")
//     return
//   }
//   w.Header().Set("Content-Type", "application/json")
//   w.WriteHeader(http.StatusOK)
//   fmt.Fprintf(w, "%s\n", resJson)
// }
