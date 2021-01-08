package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token struct {
	MeterNumber string             `json:"meterNumber"`
	TokenString string             `json:"tokenString"`
	ID          primitive.ObjectID `json:"ID"`
	UpdatedBy   *User              `json:"updatedBy"`
	CreatedBy   *User              `json:"createdBy"`
}

func (Token) IsBaseObject() {}

type TokenConnection struct {
	// A list of the meters, paginated by the provided values
	Data []*Token `json:"data"`
	// Information for paginating this connection
	PageInfo *PageInfo `json:"pageInfo"`
}

func (c *TokenConnection) CreateConection(limit int64) *TokenConnection {
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
