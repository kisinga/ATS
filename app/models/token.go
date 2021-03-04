package models

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	MeterNumber string             `json:"meterNumber,omitempty" bson:"meterNumber,omitempty"`
	TokenString string             `json:"tokenString,omitempty" bson:"tokenString,omitempty"`
	ID          primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`
	Status      TokenStatus        `json:"status,omitempty" bson:"status,omitempty"`
	APIKey      primitive.ObjectID `json:"apiKey,omitempty" bson:"apiKey,omitempty"`
}

// String returns the token string without the dashes
func (t Token) String() string {
	return strings.ReplaceAll(t.TokenString, "-", "")
}

type NewToken struct {
	MeterNumber string `json:"meterNumber,omitempty" binding:"required"`
	TokenString string `json:"tokenString,omitempty" binding:"required"`
}
type TokenStatus int

const (
	StatusNew TokenStatus = iota
	StatusSent
	StatusError
	StatusLoading
	StatusLoaded
)

func (y *TokenStatus) UnmarshalGQL(v interface{}) error {
	status, ok := v.(int)
	if !ok {
		return fmt.Errorf("YesNo must be an int")
	}
	*y = TokenStatus(status)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (y TokenStatus) MarshalGQL(w io.Writer) {
	w.Write([]byte([]byte(strconv.Itoa(int(y)))))
}

type TokenConnection struct {
	// A list of the meters, paginated by the provided values
	Data []*Token `json:"data"`
	// Information for paginating this connection
	PageInfo *PageInfo `json:"pageInfo"`
}

func (c *TokenConnection) CreateConection(limit int64, count int64) *TokenConnection {
	if len(c.Data) > 0 {
		EndCursor := primitive.ObjectID{}
		if int64(len(c.Data)) > limit-1 {
			EndCursor = c.Data[len(c.Data)-2].ID
		} else {
			EndCursor = c.Data[len(c.Data)-1].ID
		}
		c.PageInfo = &PageInfo{
			//Deduct 2 items in order to get the last
			EndCursor: EndCursor,
			HasNextPage: func() bool {
				if int64(len(c.Data)) > limit-1 {
					return true
				}
				return false
			}(),
			Count:       count,
			StartCursor: c.Data[0].ID,
		}
		// Make sure to return the specified number of elements only
		if int64(len(c.Data)) > limit-1 {
			c.Data = c.Data[:len(c.Data)-1]
		}
	} else {
		c.PageInfo = &PageInfo{}
	}
	return c
}
