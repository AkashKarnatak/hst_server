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

type UserController struct {
  startupColl *mongo.Collection
  mentorColl *mongo.Collection
}

func NewUserController(startupColl *mongo.Collection,
  mentorColl *mongo.Collection) *UserController {
  return &UserController{startupColl, mentorColl}
}

func (mc *UserController) GetMentors(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  cursor, err := mc.mentorColl.Find(ctx, bson.M{})
  if err != nil {
    log.Fatalf("Error in retrieving data\n%v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  var res []models.Mentor
  err = cursor.All(ctx, &res)
  if err != nil {
    log.Fatalf("Unable to parse collection data\n%v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  resJson, err := json.Marshal(res)
  if err != nil {
    log.Fatalf("Unable to marshal data to json\n%v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n", resJson)
}

func (mc *UserController) GetStartups(w http.ResponseWriter,
  r *http.Request, _ httprouter.Params) {
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  cursor, err := mc.startupColl.Find(ctx, bson.M{})
  if err != nil {
    log.Fatalf("Error in retrieving data\n%v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  var res []models.Startup
  err = cursor.All(ctx, &res)
  if err != nil {
    log.Fatalf("Unable to parse collection data\n%v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  resJson, err := json.Marshal(res)
  if err != nil {
    log.Fatalf("Unable to marshal data to json\n%v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n", resJson)
}
