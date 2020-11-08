package auth

import (
	"context"
	"net/http"
)

type (
	Option struct {
		Enable bool
	}
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if option.Enable {
			sessionID := r.Header.Get("session")
			if sessionID == "" {
				RespondError(w, "Unauthorized Access", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "session", sessionID)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
