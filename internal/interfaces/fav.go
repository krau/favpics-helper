package interfaces

import "github.com/krau/favpics-helper/internal/structs"

type Fav interface {
	NewFavPics() ([]structs.Pic, error)
}
