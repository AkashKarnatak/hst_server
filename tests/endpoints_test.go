package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AkashKarnatak/hst_server/controllers"
	"github.com/AkashKarnatak/hst_server/db"
	"github.com/julienschmidt/httprouter"
)

func TestGetMentors(t *testing.T) {
  // connect to database
  client := db.ConnectDB()
  defer db.DisconnectDB(client)
  mentorColl := client.Database("hst").Collection("mentors")
  mc := controllers.NewMentorController(mentorColl)
  // setup router
  router := httprouter.New()
  router.GET("/mentor", mc.GetMentors)
  req, err := http.NewRequest("GET", "/mentor", nil)
  if err != nil {
    log.Fatalln(err)
  }
  rr := httptest.NewRecorder()
  router.ServeHTTP(rr, req)
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("Wrong status. Required %v, got %v", http.StatusOK, rr.Code)
  }
}

func TestGetStartups(t *testing.T) {
  // connect to database
  client := db.ConnectDB()
  defer db.DisconnectDB(client)
  startupColl := client.Database("hst").Collection("startups")
  sc := controllers.NewStartupController(startupColl)
  // setup router
  router := httprouter.New()
  router.GET("/startup", sc.GetStartups)
  req, err := http.NewRequest("GET", "/startup", nil)
  if err != nil {
    log.Fatalln(err)
  }
  rr := httptest.NewRecorder()
  router.ServeHTTP(rr, req)
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("Wrong status. Required %v, got %v", http.StatusOK, rr.Code)
  }
}

func TestGetEvents(t *testing.T) {
  // connect to database
  client := db.ConnectDB()
  defer db.DisconnectDB(client)
  eventColl := client.Database("hst").Collection("events")
  mc := controllers.NewEventController(eventColl)
  // setup router
  router := httprouter.New()
  router.GET("/event", mc.GetEvents)
  req, err := http.NewRequest("GET", "/event", nil)
  if err != nil {
    log.Fatalln(err)
  }
  rr := httptest.NewRecorder()
  router.ServeHTTP(rr, req)
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("Wrong status. Required %v, got %v", http.StatusOK, rr.Code)
  }
}
