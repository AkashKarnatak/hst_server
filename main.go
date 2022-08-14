package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AkashKarnatak/hst_server/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
  addr := os.Getenv("ADDR")
  dbUri := os.Getenv("DB_URI")
  // using julienschmidt router
  mux := httprouter.New()
  // connect to mongodb
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
  if err != nil {
    log.Fatalf("Unable to connect to mongodb\n%v\n", err)
    return
  }
  defer func() {
    if err := client.Disconnect(ctx); err != nil {
      log.Fatalf("Unable to disconnect\n%v\n", err)
    }
  }()
  mentorColl := client.Database("hst").Collection("mentors")
  startupColl := client.Database("hst").Collection("startups")
  mc := controllers.NewMentorController(mentorColl)
  sc := controllers.NewStartupController(startupColl)
  mux.GET("/mentor", mc.GetMentors)
  mux.GET("/startup", sc.GetStartups)

  fmt.Printf("Listening on %v\n", addr)
  http.ListenAndServe(addr, mux)
}
