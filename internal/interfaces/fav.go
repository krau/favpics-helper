package interfaces

import "github.com/krau/favpics-helper/internal/models"

type Fav interface {
	NewFavPics() ([]models.Pic, error)
}
