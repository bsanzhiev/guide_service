package ports

// Этот интерфейс не содержит реализации, он лишь задает контракт.
// Реальная реализация будет в adapters, например,
// в adapters/memory/repository.go для хранения данных
// в памяти или в другом адаптере для работы с базой данных.

import "github.com/bsanzhiev/guide_service/domain"

// PlaceRepository определяет интерфейс для управления местами.
type PlaceRepository interface {
	AddPlace(p domain.Place) error

	GetPlace(id string) (domain.Place, error)

	UpdatePlace(id string, p domain.Place) error

	DeletePlace(id string) error

	GetAllPlaces() ([]domain.Place, error)
}
