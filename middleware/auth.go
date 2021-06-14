package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/moeabdol/bookstore-api-golang/utils"
)

// AuthMiddleware function
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if len(authorization) == 0 {
			utils.Log.Error("Authorization header is not provided in request")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fields := strings.Fields(authorization)
		if len(fields) < 2 {
			utils.Log.Error("Invalid authorization header format")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			utils.Log.Errorf("Unsupported authorization type %s", authorizationType)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		accessToken := fields[1]
		payload, err := utils.VerifyToken(accessToken, "keys/public-key.pem")
		if err != nil {
			utils.Log.Error(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		type key string
		const k key = "payload"
		ctx := context.WithValue(r.Context(), k, payload)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
