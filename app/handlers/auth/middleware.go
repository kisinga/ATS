package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/kisinga/ATS/app/domain"
	"github.com/kisinga/ATS/app/models"
)

//Middleware adds the user to the current context
//https://gqlgen.com/recipes/authentication/
func Middleware(firebase *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Passthrough websocket requests
		if ws := c.Request.Header.Get("Upgrade"); ws == "websocket" {
			c.Next()
			return
		}
		// Get the ID token sent by the client
		session, err := c.Cookie("session")
		token := ""
		if err != nil {
			// Session cookie is unavailable. Force user to login.
			// http.Redirect(w, r, "/login", http.StatusFound)
			// return
			token = tokenFromHeader(c.Request)
		} else {
			token = session
		}
		if token == "" {
			// Session is invalid. Force user to login.
			c.AbortWithStatusJSON(http.StatusForbidden, "Invalid authtoken")
			return
		}
		user, err := GetUserFromToken(c.Request.Context(), firebase, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		// Token is authenticated, pass it through
		// fmt.Println(token)
		// fmt.Println(user)
		// put it in context
		c.Set("user", user)
		// ctx := context.WithValue(c.Request.Context(), userCtxKey, user)
		// // // and call the next with our new context
		// c.Request = c.Request.WithContext(ctx)
		// next.ServeHTTP(w, r)
		c.Next()
	}
}

func GetUserFromToken(ctx context.Context, firebase *firebase.App, token string) (*models.User, error) {
	// Verify the session cookie. In this case an additional check is added to detect
	// if the user's Firebase session was revoked, user deleted/disabled, etc.
	client, err := firebase.Auth(context.Background())
	firebaseToken, err := client.VerifySessionCookieAndCheckRevoked(ctx, token)
	if err != nil {
		return nil, err
	}
	user := models.User{}
	dbByte, err := json.Marshal(firebaseToken.Claims)
	err = json.Unmarshal(dbByte, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
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

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
func GinContextToContext(c *gin.Context) context.Context {
	ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
	c.Request = c.Request.WithContext(ctx)
	return ctx
}
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func GetUser(ctx *gin.Context) *models.User {
	v, exists := ctx.Get("user")
	if !exists {
		return nil
	}
	vv := v.(*models.User)
	return vv
}

func GetUserFromContext(ctx context.Context, domain *domain.Domain) (*models.User, error) {
	cc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// me := ForContext(ctx)
	me := GetUser(cc)
	if me == nil {
		return nil, errors.New("failed extracting user from context")
	}
	// user in context doesnt have ID field
	me, err = domain.User.GetUser(ctx, me.Email)
	if err != nil {
		return nil, err
	}
	return me, nil
}
