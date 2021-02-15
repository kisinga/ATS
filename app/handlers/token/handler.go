package handlers

import (
	"net/http"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/kisinga/ATS/app/models"
	"github.com/kisinga/ATS/app/registry"
)

// TokenRegex is used to validate the format of a token string
var TokenRegex = regexp.MustCompile(`^(?:\w{4}-){4}\w{4}$`)

func TokenHandler(domain *registry.Domain) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawKey := c.Request.Header.Get("Authorization")
		if len(rawKey) == 0 {
			c.AbortWithStatusJSON(401, "Missing API Key in Header")
			return
		}
		key, err := primitive.ObjectIDFromHex(rawKey)
		if err != nil {
			c.AbortWithStatusJSON(401, "Invalid API Key in Header")
			return
		}
		if rawKey != domain.APIKey.GetLatest().ID.Hex() {
			c.AbortWithStatusJSON(401, "Invalid API Key in Header")
			return
		}
		var token models.NewToken
		if err := c.ShouldBind(&token); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid request")
			return
		}
		if !TokenRegex.Match([]byte(token.TokenString)) {
			c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid token")
			return
		}

		meter, err := domain.Meter.GetMeter(c.Request.Context(), token.MeterNumber)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, "Invalid meter number")
			return
		}
		if !meter.Active {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, "Meter not active")
			return
		}
		_, err = domain.Token.AddToken(c.Request.Context(), token, key)
		if err != nil {
			c.AbortWithStatusJSON(500, err)
			return
		}

		c.JSON(200, "success")
	}
}
