package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AkashKarnatak/hst_server/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserClaim struct {
	jwt.RegisteredClaims
	Id string
}

func GenerateNewToken(id string) (string, error) {
  key := []byte(os.Getenv("JWT_KEY"))
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
    RegisteredClaims: jwt.RegisteredClaims{ IssuedAt: jwt.NewNumericDate(time.Now()) },
    Id: id,
  })
  s, err := token.SignedString(key)
  if err != nil {
    return "", err
  }
  return s, nil
}

func ParseTokenStr(tokenString string) (UserClaim, error) {
  var uc UserClaim
  token, err := jwt.ParseWithClaims(tokenString, &uc, func(*jwt.Token) (interface{}, error) {
    return []byte(os.Getenv("JWT_KEY")), nil
  })
  if err != nil {
    return uc, fmt.Errorf("Unable to parse token string: %v", err)
  }
  if !token.Valid {
    return uc, fmt.Errorf("Invalid token")
  }
  return uc, nil
}

// TODO: find a better way to get collections
func Authorize(h httprouter.Handle,
  startupColl *mongo.Collection,
  mentorColl *mongo.Collection) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter,
		r *http.Request, p httprouter.Params) {
		tokenString := r.Header.Get("Authorization")
    uc, err := ParseTokenStr(tokenString)
    if err != nil {
      log.Printf("%v\n", err)
      w.WriteHeader(http.StatusUnauthorized)
      fmt.Fprintln(w, "Unauthorized")
      return
    }
    // identify whether user is mentor or startup
    var coll *mongo.Collection
    if uc.Id[:3] == "STA" {
      coll = startupColl
    } else {
      coll = mentorColl
    }
    // fetch the user from their collection
    ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
    defer cancel()
    var user models.User
    err = coll.FindOne(ctx, bson.M{"_id": uc.Id}).Decode(&user)
    if err != nil {
      log.Printf("Error fetching user from db: %v\n", err)
      w.WriteHeader(http.StatusUnauthorized)
      fmt.Fprintln(w, "Unauthorized")
      return
    }
    found := false
    for _, t := range user.Tokens {
      if t == tokenString {
        found = true
        break
      }
    }
    if !found {
      log.Printf("Invalid token\n")
      w.WriteHeader(http.StatusUnauthorized)
      fmt.Fprintln(w, "Unauthorized")
      return
    }
    h(w, r, httprouter.Params{{Key: "id", Value: uc.Id}})
	})
}
