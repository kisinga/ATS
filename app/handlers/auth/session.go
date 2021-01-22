package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/kisinga/ATS/app/registry"
)

// Token ...
type Token struct {
	Token  string    `json:"Bearer"`
	Expiry time.Time `json:"expiry"`
}

type data struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// SessionInit validates the provided userID using the firebase Admin API, then sets a cookie that is used to validate every request
func SessionInit(firebase *firebase.App, domain *registry.Domain) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the ID token sent by the client
		defer c.Request.Body.Close()
		claims, err := getIDTokenFromBody(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		// Set session expiration to 4 days.
		expiresIn := time.Hour * 24 * 4

		// Create the session cookie. This will also verify the ID token in the process.
		// The session cookie will have the same claims as the ID token.
		// To only allow session cookie setting on recent sign-in, auth_time in ID token
		// can be checked to ensure user was recently signed in before creating a session cookie.
		client, err := firebase.Auth(context.Background())
		cookie, err := client.SessionCookie(c.Request.Context(), claims.ID, expiresIn)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to create a session cookie")
			return
		}
		user, err := domain.User.GetUser(c.Request.Context(), claims.Email)

		if err != nil || !user.Active {
			c.AbortWithStatusJSON(http.StatusForbidden, "Not authorised")
			return
		}
		// Set cookie policy for session cookie.
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "session",
			Value:    cookie,
			MaxAge:   int(expiresIn.Seconds()),
			HttpOnly: true,
			Secure:   true,
		})
		response := Token{cookie, time.Now().Add(expiresIn).UTC()}
		JSONResponse(response, c.Writer)
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

// SessionTerm deletes the set cookie
func SessionTerm() gin.HandlerFunc {
	return func(c *gin.Context) {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: 0,
		})
		http.Redirect(c.Writer, c.Request, "/login", http.StatusFound)
	}
}

func getIDTokenFromBody(r *http.Request) (data, error) {
	claims := data{}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&claims)
	return claims, err
}
