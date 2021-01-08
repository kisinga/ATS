package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BaseModel struct {
	ID        primitive.ObjectID `json:"ID"  bson:"_id,omitempty"`
	CreatedBy primitive.ObjectID `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	UpdatedBy primitive.ObjectID `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
}
