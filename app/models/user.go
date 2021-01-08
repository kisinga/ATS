package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID    primitive.ObjectID `json:"ID,omitempty" bson:"_id,omitempty"`
	Email string             `json:"email"`
	Name  string             `json:"name"`
}
