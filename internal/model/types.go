package model

type Filter struct {
	Q           string
	Page        int
	Count       int
	WithDeleted bool
}
