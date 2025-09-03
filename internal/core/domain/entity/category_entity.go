package entity

type CategoryEntity struct {
	ID  string
	Title string
	Slug  string
	User  UserEntity
}