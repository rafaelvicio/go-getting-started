package middleware

import (
	"fmt"
	"net/http"
	"pipenator-backend/securirty"
)

type Adapter func(http.Handler) http.Handler

func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func SetHeadersDefault(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")

		next.ServeHTTP(w, r)
	})
}

func VerifyToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("token")

		temToken := securirty.Verificar(token)

		if temToken == true {
			fmt.Println("Token valido")
		} else {
			fmt.Println("Token não valido")
		}

		fmt.Println("Token", token)

		h.ServeHTTP(w, r)
	})
}
