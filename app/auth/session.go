package auth

import (
	"net/http"

	firebase "firebase.google.com/go"
)

func SessionInit(firebase *firebase.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func SessionTerm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
