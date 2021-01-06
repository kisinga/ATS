package auth

import (
	"encoding/json"
	"net/http"

	firebase "firebase.google.com/go"

	"github.com/go-chi/jwtauth"
)

// Token ...
type Token struct {
	Token string `json:"Bearer"`
}

// SessionInit validates the provided userID using the firebase Admin API, then sets a cookie that is used to validate every request
func SessionInit(firebase *firebase.App, tokenAuth *jwtauth.JWTAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

		// For debugging/example purposes, we generate and print
		// a sample jwt token with claims `user_id:123` here:
		_, tokenString, err := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
		if err != nil {
			http.Error(w, "Error parsing jwt", http.StatusInternalServerError)
		}
		response := Token{tokenString}
		JSONResponse(response, w)
	}
}

// SessionTerm deletes the set cookie
func SessionTerm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// JSONResponse ...
func JSONResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
