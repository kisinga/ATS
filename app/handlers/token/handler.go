package handlers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/kisinga/ATS/app/models"
	"github.com/kisinga/ATS/app/registry"
)

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
			c.AbortWithStatusJSON(401, "Invalid request")
			return
		}
		_, err = domain.Token.AddToken(c.Request.Context(), token, key)
		if err != nil {
			c.AbortWithStatusJSON(500, "Error Adding token to DB")
			return
		}

		c.JSON(200, "success")
	}
}
