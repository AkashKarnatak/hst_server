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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventController struct {
  coll *mongo.Collection
}

func NewEventController(coll *mongo.Collection) *EventController {
  return &EventController{coll}
}

func (mc *EventController) GetEvents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  opts := options.Find().SetSort(bson.M{"index": 1})
  cursor, err := mc.coll.Find(ctx, bson.M{}, opts)
  if err != nil {
    log.Printf("Error in retrieving data: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  var res []models.Event
  err = cursor.All(ctx, &res)
  if err != nil {
    log.Printf("Unable to parse collection data: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  resJson, err := json.Marshal(res)
  if err != nil {
    log.Printf("Unable to marshal data to json: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n", resJson)
}
