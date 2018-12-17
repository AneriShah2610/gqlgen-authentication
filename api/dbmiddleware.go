package api

import (
	"context"
	"net/http"

	"github.com/aneri/gqlgen-authentication/dal"
)

var ctxt context.Context

func MiddleWareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		crConn := dal.Connect()
		ctxt = context.WithValue(request.Context(), "crConn", crConn)
		next.ServeHTTP(writer, request.WithContext(ctxt))
	})
}
