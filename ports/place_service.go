package ports

import "github.com/bsanzhiev/guide_service/domain"

type PlaceService interface {
	CreatePlace(p domain.Place) (domain.Place, error)

	GetPlaceByID(id string) (domain.Place, error)

	UpdatePlace(id string, p domain.Place) (domain.Place, error)

	DeletePlace(id string) error

	GetAllPlaces() ([]domain.Place, error)
}
