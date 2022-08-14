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
  mc := controllers.NewMentorController(mentorColl)
  sc := controllers.NewStartupController(startupColl)
  // setup router
  router := httprouter.New()
  router.GET("/mentor", mc.GetMentors)
  router.GET("/startup", sc.GetStartups)

  fmt.Printf("Listening on %v\n", addr)
  http.ListenAndServe(addr, router)
}
