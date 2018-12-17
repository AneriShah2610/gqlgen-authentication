package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/aneri/gqlgen-authentication/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var contxt context.Context
var secretkey = []byte("secret_key")
var user models.User

func AuthHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		header := request.Header.Get("Authorization")
		if header == "" {
			next.ServeHTTP(response, request)
		} else {
			token := jwt.New(jwt.SigningMethodHS256)
			token.Claims = jwt.MapClaims{
				"userid":   user.ID,
				"username": user.Name,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			}
			tokenstring, err := token.SignedString(secretkey)
			if err != nil {
				log.Fatal("Error while generating token ", err)
			}
			ctxt := context.WithValue(request.Context(), "User", tokenstring)
			next.ServeHTTP(response, request.WithContext(ctxt))
		}
	})
}
