package middleware

import "net/http"

type Middleware struct {
	LoggingMW func(next http.Handler) http.Handler
}
