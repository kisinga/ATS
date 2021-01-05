package models

type User struct {
	ID    string `json:"ID" bson:"_id,omitempty"`
	Email string `json:"email"`
}
