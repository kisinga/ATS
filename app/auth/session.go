package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
)

// Token ...
type Token struct {
	Token string `json:"Bearer"`
}

// SessionInit validates the provided userID using the firebase Admin API, then sets a cookie that is used to validate every request
func SessionInit(firebase *firebase.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the ID token sent by the client
		defer r.Body.Close()
		idToken, err := getIDTokenFromBody(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Set session expiration to 4 days.
		expiresIn := time.Hour * 24 * 4

		// Create the session cookie. This will also verify the ID token in the process.
		// The session cookie will have the same claims as the ID token.
		// To only allow session cookie setting on recent sign-in, auth_time in ID token
		// can be checked to ensure user was recently signed in before creating a session cookie.
		client, err := firebase.Auth(context.Background())
		cookie, err := client.SessionCookie(r.Context(), idToken, expiresIn)
		if err != nil {
			http.Error(w, "Failed to create a session cookie", http.StatusInternalServerError)
			return
		}

		// Set cookie policy for session cookie.
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    cookie,
			MaxAge:   int(expiresIn.Seconds()),
			HttpOnly: true,
			Secure:   true,
		})
		w.Write([]byte(`{"status": "success"}`))
	}
}

// SessionTerm deletes the set cookie
func SessionTerm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: 0,
		})
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func getIDTokenFromBody(r *http.Request) (string, error) {
	cookie := struct {
		ID string `json:"id,omitempty"`
	}{}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&cookie)
	return cookie.ID, err
}
