package models

type User struct {
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	*BaseModel
	Disabled bool `json:"disabled,omitempty" bson:"disabled,omitempty"`
}

func (User) IsBaseObject() {}
