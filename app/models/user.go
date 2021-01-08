package models

type User struct {
	Email string `json:"email,omitempty" bson:"email"`
	Name  string `json:"name,omitempty" bson:"name"`
	BaseModel
	Disabled bool `json:"disabled,omitempty" bson:"disabled"`
}

func (User) IsBaseObject() {}
