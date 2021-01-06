package models

type User struct {
	ID    string `json:"ID,omitempty" bson:"_id,omitempty"`
	Email string `json:"email"`
}
