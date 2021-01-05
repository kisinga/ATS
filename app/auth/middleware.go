package auth

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// A stand-in for our database backed user object
type User struct {
	Name    string
	IsAdmin bool
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(firebase *firebase.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth-cookie")

			// reject unauthenticated requests
			if err != nil || c == nil {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			// if !domain.User.ValidLogin(c) {
			// 	http.Error(w, "Invalid cookie", http.StatusForbidden)
			// 	return
			// }
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}
