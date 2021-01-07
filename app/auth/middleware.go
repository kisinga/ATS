package auth

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
)

func Middleware(firebase *firebase.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the ID token sent by the client
			session, err := r.Cookie("session")
			token := ""
			if err != nil {
				// Session cookie is unavailable. Force user to login.
				// http.Redirect(w, r, "/login", http.StatusFound)
				// return
				token = tokenFromHeader(r)
			} else {
				token = session.Value
			}
			if token == "" {
				// Session is invalid. Force user to login.
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			// Verify the session cookie. In this case an additional check is added to detect
			// if the user's Firebase session was revoked, user deleted/disabled, etc.
			client, err := firebase.Auth(context.Background())
			_, err = client.VerifySessionCookieAndCheckRevoked(r.Context(), token)
			if err != nil {
				// Session cookie is invalid. Force user to login.
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			// Token is authenticated, pass it through
			next.ServeHTTP(w, r)
		})
	}
}

// TokenFromHeader tries to retreive the token string from the
// "Authorization" reqeust header: "Authorization: BEARER T".
func tokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
