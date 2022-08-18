package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
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
    return uc, fmt.Errorf("Unable to parse token string")
  }
  if !token.Valid {
    return uc, fmt.Errorf("Invalid token")
  }
  return uc, nil
}

func Authorize(h httprouter.Handle) httprouter.Handle {
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
    h(w, r, httprouter.Params{{Key: "id", Value: uc.Id}})
	})
}
