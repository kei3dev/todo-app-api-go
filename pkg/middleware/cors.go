package middleware

import (
	"net/http"
	"os"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		backendOrigin := os.Getenv("BACKEND_ORIGIN")
		frontendOrigin := os.Getenv("FRONTEND_ORIGIN")

		origin := r.Header.Get("Origin")
		if origin != "" {
			if origin == backendOrigin || origin == frontendOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Max-Age", "300") // Set preflight request cache time to 5 minutes
			} else {
				http.Error(w, "Not allowed by CORS", http.StatusForbidden)
				return
			}
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
