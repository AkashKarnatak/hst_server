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
  themeMappingColl := client.Database("hst").Collection("themeMappings")
  mappingColl := client.Database("hst").Collection("mappings")
  guestColl := client.Database("hst").Collection("guests")
  uc := controllers.NewUserController(startupColl, mentorColl, guestColl)
  ec := controllers.NewEventController(eventColl)
  mc := controllers.NewMeetingController(mappingColl, themeMappingColl)
  // setup router
  router := httprouter.New()
  router.GET("/mentor", uc.GetMentors)
  router.GET("/incubatedStartup", uc.GetIncubatedStartups)
  router.GET("/hstStartup", uc.GetHstStartups)
  router.GET("/event", ec.GetEvents)
  router.POST("/login", uc.Login)
  // TODO: find a better way to pass collection to authorization middleware
  router.POST("/logout", middlewares.Authorize(uc.Logout, startupColl, mentorColl))
  router.POST("/logoutAll", middlewares.Authorize(uc.LogoutAll, startupColl, mentorColl))
  router.POST("/check", middlewares.Authorize(uc.Check, startupColl, mentorColl))
  router.POST("/meeting", middlewares.Authorize(mc.GetMeetings, startupColl, mentorColl))
  router.POST("/guest", uc.CreateGuest)

  fmt.Printf("Listening on %v\n", addr)
  http.ListenAndServe(addr, router)
}
