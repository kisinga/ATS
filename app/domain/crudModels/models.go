package crudModels

// RepoOperation defines a int representation of the operation taking place in each domain
type RepoOperation int

const (
	Create RepoOperation = iota
	Read
	Update
	Delete
)

type DomainName int

const (
	Ingredient DomainName = iota
	Recipe
	User
)
