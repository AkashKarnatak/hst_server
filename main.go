package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AkashKarnatak/hst_server/controllers"
	"github.com/AkashKarnatak/hst_server/db"
	"github.com/AkashKarnatak/hst_server/middlewares"
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
  mappingColl := client.Database("hst").Collection("mappings")
  uc := controllers.NewUserController(startupColl, mentorColl)
  ec := controllers.NewEventController(eventColl)
  mc := controllers.NewMeetingController(mappingColl)
  // setup router
  router := httprouter.New()
  router.GET("/mentor", uc.GetMentors)
  router.GET("/startup", uc.GetStartups)
  router.GET("/event", ec.GetEvents)
  router.POST("/login", uc.Login)
  router.POST("/logout", middlewares.Authorize(uc.Logout))
  router.POST("/logoutAll", middlewares.Authorize(uc.LogoutAll))
  router.POST("/meeting", middlewares.Authorize(mc.GetMeetings))

  fmt.Printf("Listening on %v\n", addr)
  http.ListenAndServe(addr, router)
}
