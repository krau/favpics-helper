package models

type Source interface {
	NewFavPics() ([]Pic, error)
}
