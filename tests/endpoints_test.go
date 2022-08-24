package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/AkashKarnatak/hst_server/controllers"
	"github.com/AkashKarnatak/hst_server/db"
	"github.com/AkashKarnatak/hst_server/middlewares"
	"github.com/julienschmidt/httprouter"
)

var tokenString string

func TestGetMentors(t *testing.T) {
  // connect to database
  client := db.ConnectDB()
  defer db.DisconnectDB(client)
  startupColl := client.Database("hst").Collection("startups")
  mentorColl := client.Database("hst").Collection("mentors")
  guestColl := client.Database("hst").Collection("guests")
  uc := controllers.NewUserController(startupColl, mentorColl, guestColl)
  // setup router
  router := httprouter.New()
  router.GET("/mentor", uc.GetMentors)
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

func TestGetIncubatedStartups(t *testing.T) {
  // connect to database
  client := db.ConnectDB()
  defer db.DisconnectDB(client)
  startupColl := client.Database("hst").Collection("startups")
  mentorColl := client.Database("hst").Collection("mentors")
  guestColl := client.Database("hst").Collection("guests")
  uc := controllers.NewUserController(startupColl, mentorColl, guestColl)
  // setup router
  router := httprouter.New()
  router.GET("/incubatedStartup", uc.GetIncubatedStartups)
  req, err := http.NewRequest("GET", "/incubatedStartup", nil)
  if err != nil {
    log.Fatalln(err)
  }
  rr := httptest.NewRecorder()
  router.ServeHTTP(rr, req)
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("Wrong status. Required %v, got %v", http.StatusOK, rr.Code)
  }
}

func TestGetHstStartups(t *testing.T) {
  // connect to database
  client := db.ConnectDB()
  defer db.DisconnectDB(client)
  startupColl := client.Database("hst").Collection("startups")
  mentorColl := client.Database("hst").Collection("mentors")
  guestColl := client.Database("hst").Collection("guests")
  uc := controllers.NewUserController(startupColl, mentorColl, guestColl)
  // setup router
  router := httprouter.New()
  router.GET("/hstStartup", uc.GetHstStartups)
  req, err := http.NewRequest("GET", "/hstStartup", nil)
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

func TestLogin(t *testing.T) {
  // connect to database
  client := db.ConnectDB()
  defer db.DisconnectDB(client)
  startupColl := client.Database("hst").Collection("startups")
  mentorColl := client.Database("hst").Collection("mentors")
  guestColl := client.Database("hst").Collection("guests")
  uc := controllers.NewUserController(startupColl, mentorColl, guestColl)
  // setup router
  router := httprouter.New()
  router.POST("/login", uc.Login)
  // create new request
  req, err := http.NewRequest("POST", "/login", nil)
  form := url.Values{
    "email": {"siddharth@thesolarlabs.com"},
    "phNo": {"8628868995"},
  }
  req.PostForm = form
  if err != nil {
    log.Fatalln(err)
  }
  rr := httptest.NewRecorder()
  router.ServeHTTP(rr, req)
  var auth map[string]string
  err = json.NewDecoder(rr.Body).Decode(&auth)
  if err != nil {
    log.Fatalln(err)
  }
  tokenString = auth["token"]
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("Wrong status. Required %v, got %v", http.StatusOK, rr.Code)
  }
}

func TestGetMeetings(t *testing.T) {
  // connect to database
  client := db.ConnectDB()
  defer db.DisconnectDB(client)
  mappingColl := client.Database("hst").Collection("mappings")
  startupColl := client.Database("hst").Collection("startups")
  mentorColl := client.Database("hst").Collection("mentors")
  mc := controllers.NewMeetingController(mappingColl)
  // setup router
  router := httprouter.New()
  router.POST("/meeting", middlewares.Authorize(mc.GetMeetings, startupColl, mentorColl))
  // create new request
  req, err := http.NewRequest("POST", "/meeting", nil)
  req.Header.Set("Authorization", tokenString)
  if err != nil {
    log.Fatalln(err)
  }
  rr := httptest.NewRecorder()
  router.ServeHTTP(rr, req)
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("Wrong status. Required %v, got %v", http.StatusOK, rr.Code)
  }
}

func TestLogout(t *testing.T) {
  // connect to database
  client := db.ConnectDB()
  defer db.DisconnectDB(client)
  startupColl := client.Database("hst").Collection("startups")
  mentorColl := client.Database("hst").Collection("mentors")
  guestColl := client.Database("hst").Collection("guests")
  uc := controllers.NewUserController(startupColl, mentorColl, guestColl)
  // setup router
  router := httprouter.New()
  router.POST("/logout", middlewares.Authorize(uc.Logout, startupColl, mentorColl))
  // create new request
  req, err := http.NewRequest("POST", "/logout", nil)
  req.Header.Set("Authorization", tokenString)
  if err != nil {
    log.Fatalln(err)
  }
  rr := httptest.NewRecorder()
  router.ServeHTTP(rr, req)
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("Wrong status. Required %v, got %v", http.StatusOK, rr.Code)
  }
}

func TestLogoutAll(t *testing.T) {
  // login first
  TestLogin(t)
  // connect to database
  client := db.ConnectDB()
  defer db.DisconnectDB(client)
  startupColl := client.Database("hst").Collection("startups")
  mentorColl := client.Database("hst").Collection("mentors")
  guestColl := client.Database("hst").Collection("guests")
  uc := controllers.NewUserController(startupColl, mentorColl, guestColl)
  // setup router
  router := httprouter.New()
  router.POST("/logoutAll", middlewares.Authorize(uc.LogoutAll, startupColl, mentorColl))
  // create new request
  req, err := http.NewRequest("POST", "/logoutAll", nil)
  req.Header.Set("Authorization", tokenString)
  if err != nil {
    log.Fatalln(err)
  }
  rr := httptest.NewRecorder()
  router.ServeHTTP(rr, req)
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("Wrong status. Required %v, got %v", http.StatusOK, rr.Code)
  }
}
