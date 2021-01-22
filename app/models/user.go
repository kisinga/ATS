package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	ID        primitive.ObjectID `json:"ID"  bson:"_id,omitempty"`
	CreatedBy primitive.ObjectID `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	UpdatedBy primitive.ObjectID `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
	Active    bool               `json:"active" bson:"active"`
}

func (User) IsBaseObject() {}

type UsersConnection struct {
	// A list of the meters, paginated by the provided values
	Data []*User `json:"data"`
	// Information for paginating this connection
	PageInfo *PageInfo `json:"pageInfo"`
}

func (c *UsersConnection) CreateConection(limit int64) *UsersConnection {
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
