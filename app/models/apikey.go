package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type APIKey struct {
	ID        primitive.ObjectID `json:"ID"  bson:"_id,omitempty"`
	CreatedBy primitive.ObjectID `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
}

func (APIKey) IsBaseObject() {}
