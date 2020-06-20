package entity

type User struct {
	Username      string
	FullName      string
	Password      string
	FavCategories []string
	FavSources    []string
}
