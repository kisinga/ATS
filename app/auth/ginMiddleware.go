package auth

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func Mid(firebase *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// // Get the ID token sent by the client
		// session, err := c.Cookie("session")
		// token := ""
		// if err != nil {
		// 	// Session cookie is unavailable. Force user to login.
		// 	// http.Redirect(w, r, "/login", http.StatusFound)
		// 	// return
		// 	token = tokenFromHeader(c.Request)
		// } else {
		// 	token = session
		// }
		// if token == "" {
		// 	json, _ := json.Marshal("Invalid token")
		// 	// Session is invalid. Force user to login.
		// 	c.AbortWithStatusJSON(http.StatusForbidden, json)
		// 	return
		// }
		// // Verify the session cookie. In this case an additional check is added to detect
		// // if the user's Firebase session was revoked, user deleted/disabled, etc.
		// client, err := firebase.Auth(context.Background())
		// firebaseToken, err := client.VerifySessionCookieAndCheckRevoked(c.Request.Context(), token)
		// if err != nil {
		// 	// Session cookie is invalid. Force user to login.
		// 	json, _ := json.Marshal("Invalid token")
		// 	// Session is invalid. Force user to login.
		// 	c.AbortWithStatusJSON(http.StatusForbidden, json)
		// 	return
		// }
		// user := models.User{}
		// dbByte, err := json.Marshal(firebaseToken.Claims)
		// err = json.Unmarshal(dbByte, &user)
		// if err != nil {
		// 	// Session cookie is invalid. Force user to login.
		// 	json, _ := json.Marshal("Error Marshalling userdata into json")
		// 	// Session is invalid. Force user to login.
		// 	c.AbortWithStatusJSON(http.StatusInternalServerError, json)
		// 	return
		// }
		// // Token is authenticated, pass it through
		// fmt.Println(user)
		// // put it in context
		// // ctx := context.WithValue(r.Context(), userCtxKey, user)

		// // // and call the next with our new context
		// r = r.WithContext(ctx)
		// next.ServeHTTP(w, r)
		c.Next()
	}
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
