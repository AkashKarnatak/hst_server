package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AkashKarnatak/hst_server/middlewares"
	"github.com/AkashKarnatak/hst_server/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserController struct {
  startupColl *mongo.Collection
  mentorColl *mongo.Collection
  guestColl *mongo.Collection
}

func NewUserController(startupColl *mongo.Collection,
  mentorColl *mongo.Collection,
  guestColl *mongo.Collection) *UserController {
  return &UserController{startupColl, mentorColl, guestColl}
}

func (uc *UserController) GetMentors(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  opts := options.Find().SetProjection(bson.M{
    "_id": 0,
    "Email": 0,
    "Phone": 0,
    "tokens": 0,
  })
  cursor, err := uc.mentorColl.Find(ctx, bson.M{}, opts)
  if err != nil {
    log.Printf("Error in retrieving data: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  var res []models.Mentor
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

func (uc *UserController) GetStartups(w http.ResponseWriter,
  r *http.Request, _ httprouter.Params) {
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  opts := options.Find().SetProjection(bson.M{
    "_id": 0,
    "Email": 0,
    "Phone": 0,
    "tokens": 0,
  })
  cursor, err := uc.startupColl.Find(ctx, bson.M{}, opts)
  if err != nil {
    log.Printf("Error in retrieving data: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  var res []models.Startup
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

func (uc *UserController) Login(w http.ResponseWriter,
  r *http.Request, _ httprouter.Params) {
  // get user with email and phNo in db
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  var user models.User
  coll := uc.startupColl
  err := coll.FindOne(
    ctx,
    bson.M{"Email": r.FormValue("email"), "Phone": r.FormValue("phNo")},
  ).Decode(&user)
  if err != nil {
    user = models.User{}
    coll := uc.mentorColl
    err := coll.FindOne(
      ctx,
      bson.M{"Email": r.FormValue("email"), "Phone": r.FormValue("phNo")},
    ).Decode(&user)
    if err != nil {
      log.Printf("Unable to find user: %v\n", err)
      w.WriteHeader(http.StatusNotFound)
      fmt.Fprintln(w, "User not found")
      return
    }
  }
  // now that we have a user, create jwt token for them
  tokenString, err := middlewares.GenerateNewToken(user.Id) 
  if err != nil {
    log.Printf("Unable to create new token: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  user.Tokens = append(user.Tokens, tokenString)
  // update user entry in db
  _, err = coll.UpdateByID(ctx, user.Id, bson.M{
    "$set": user,
  })
  if err != nil {
    log.Printf("Unable to udpate user: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  // return this token as json
  js, err := json.Marshal(map[string]string{"token": tokenString})
  if err != nil {
    log.Printf("Unable to marshal token to json: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n", js)
}

func (uc *UserController) Check(w http.ResponseWriter,
  r *http.Request, p httprouter.Params) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "Authorized")
}

func (uc *UserController) Logout(w http.ResponseWriter,
  r *http.Request, p httprouter.Params) {
  // identify whether user is mentor or startup
  var coll *mongo.Collection
  if p.ByName("id")[:3] == "STA" {
    coll = uc.startupColl
  } else {
    coll = uc.mentorColl
  }
  // fetch the user from their collection
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  var user models.User
  err := coll.FindOne(ctx, bson.M{"_id": p.ByName("id")}).Decode(&user)
  if err != nil {
    log.Printf("Error fetching user from db: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  // delete user's current token
  var ts []string
  authToken := r.Header.Get("Authorization") 
  for _, t := range user.Tokens {
    if t != authToken {
      ts = append(ts, t)
    }
  }
  user.Tokens = ts
  // update user entry in db
  _, err = coll.UpdateByID(ctx, user.Id, bson.M{
    "$set": user,
  })
  if err != nil {
    log.Printf("Unable to udpate user: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "Logged out successfully")
}

func (uc *UserController) LogoutAll(w http.ResponseWriter,
  r *http.Request, p httprouter.Params) {
  // identify whether user is mentor or startup
  var coll *mongo.Collection
  if p.ByName("id")[:3] == "STA" {
    coll = uc.startupColl
  } else {
    coll = uc.mentorColl
  }
  // fetch the user from their collection
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  var user models.User
  err := coll.FindOne(ctx, bson.M{"_id": p.ByName("id")}).Decode(&user)
  if err != nil {
    log.Printf("Error fetching user from db: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  // delete all user tokens
  user.Tokens = []string{}
  // update user entry in db
  _, err = coll.UpdateByID(ctx, user.Id, bson.M{
    "$set": user,
  })
  if err != nil {
    log.Printf("Unable to udpate user: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "Logged out successfully from all devices")
}

func (uc *UserController) CreateGuest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  guest := models.Guest{
    Id: primitive.NewObjectID(),
    EmailId: r.FormValue("email"),
  }
  _, err := uc.guestColl.InsertOne(ctx, guest)
  if err != nil {
    log.Printf("Error inserting guest in db: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Internal server error")
    return
  }
  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "Guest created successfully")
}
