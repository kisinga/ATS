package auth

import (
	"net/http"

	firebase "firebase.google.com/go"
)

// SessionInit validates the provided userID using the firebase Admin API, then sets a cookie that is used to validate every request
func SessionInit(firebase *firebase.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// SessionTerm deletes the set cookie
func SessionTerm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
