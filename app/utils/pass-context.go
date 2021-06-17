package utils

import (
	"context"
	"net/http"

	"github.com/graphql-go/handler"
)

//PassContext function
func PassContext(hnd *handler.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Credentials", "true")
		res.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, x-token")

		token := req.Header.Get("x-token")

		ctx := context.WithValue(context.Background(), "token", token)
		hnd.ContextHandler(ctx, res, req)
	})
}
