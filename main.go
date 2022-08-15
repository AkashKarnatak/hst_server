package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AkashKarnatak/hst_server/controllers"
	"github.com/AkashKarnatak/hst_server/db"
	"github.com/julienschmidt/httprouter"
)

func main() {
  addr := os.Getenv("ADDR")
  // connect to mongodb
  client := db.ConnectDB()
  defer db.DisconnectDB(client)
  mentorColl := client.Database("hst").Collection("mentors")
  startupColl := client.Database("hst").Collection("startups")
  eventColl := client.Database("hst").Collection("events")
  mc := controllers.NewMentorController(mentorColl)
  sc := controllers.NewStartupController(startupColl)
  ec := controllers.NewEventController(eventColl)
  // setup router
  router := httprouter.New()
  router.GET("/mentor", mc.GetMentors)
  router.GET("/startup", sc.GetStartups)
  router.GET("/event", ec.GetEvents)

  fmt.Printf("Listening on %v\n", addr)
  http.ListenAndServe(addr, router)
}
