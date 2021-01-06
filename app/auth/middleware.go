package auth

import (
	"context"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
)

func Middleware(firebase *firebase.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the ID token sent by the client
			cookie, err := r.Cookie("session")
			if err != nil {
				// Session cookie is unavailable. Force user to login.
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			// Verify the session cookie. In this case an additional check is added to detect
			// if the user's Firebase session was revoked, user deleted/disabled, etc.
			client, err := firebase.Auth(context.Background())
			k, err := client.VerifySessionCookieAndCheckRevoked(r.Context(), cookie.Value)
			if err != nil {
				// Session cookie is invalid. Force user to login.
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
			fmt.Println(k)
			// Token is authenticated, pass it through
			next.ServeHTTP(w, r)
		})
	}
}
