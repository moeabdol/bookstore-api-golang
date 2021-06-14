package middleware

import (
	"net/http"

	"github.com/moeabdol/bookstore-api-golang/utils"
)

// RequestLoggerMiddleware function
func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Log.Debugf("%s %s - {%s}", r.Method, r.URL, r.URL.RawQuery)

		next.ServeHTTP(w, r)
	})
}
