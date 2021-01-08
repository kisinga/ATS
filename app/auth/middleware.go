package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/kisinga/ATS/app/models"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

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
			firebaseToken, err := client.VerifySessionCookieAndCheckRevoked(r.Context(), token)
			if err != nil {
				// Session cookie is invalid. Force user to login.
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			user := models.User{}
			dbByte, err := json.Marshal(firebaseToken.Claims)
			err = json.Unmarshal(dbByte, &user)
			if err != nil {
				// Session cookie is invalid. Force user to login.
				http.Error(w, "Error Marshalling userdata into json", http.StatusInternalServerError)
				return
			}
			// Token is authenticated, pass it through
			fmt.Println(user)
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
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
