package memory

import (
	"errors"
	"sync"

	"github.com/bsanzhiev/guide_service/domain"
	"github.com/bsanzhiev/guide_service/ports"
)

// PlaceRepositoryMemory - реализация интерфейса PlaceRepository для хранения в памяти
type PlaceRepositoryMemory struct {
	places map[string]domain.Place
	mutex  sync.Mutex
}

// NewPlaceRepositoryMemory создает новый экземпляр PlaceRepositoryMemory
func NewPlaceRepositoryMemory() ports.PlaceRepository {
	return &PlaceRepositoryMemory{
		places: make(map[string]domain.Place),
	}
}

// AddPlace добавляет новое место в память
func (r *PlaceRepositoryMemory) AddPlace(p domain.Place) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.places[p.ID]; exists {
		return errors.New("place with this ID already exists")
	}
	r.places[p.ID] = p
	return nil
}

// GetPlace получает место по ID
func (r *PlaceRepositoryMemory) GetPlace(id string) (domain.Place, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	place, exists := r.places[id]
	if !exists {
		return domain.Place{}, errors.New("place not found")
	}
	return place, nil
}

// UpdatePlace обновляет существующее место в памяти
func (r *PlaceRepositoryMemory) UpdatePlace(id string, updatedPlace domain.Place) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.places[id]; !exists {
		return errors.New("place not found")
	}
	updatedPlace.ID = id // Убеждаемся, что ID обновленного места совпадает
	r.places[id] = updatedPlace
	return nil
}

// DeletePlace удаляет место из памяти по ID
func (r *PlaceRepositoryMemory) DeletePlace(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.places[id]; !exists {
		return errors.New("place not found")
	}
	delete(r.places, id)
	return nil
}

// GetAllPlaces возвращает все места из памяти
func (r *PlaceRepositoryMemory) GetAllPlaces() ([]domain.Place, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	places := make([]domain.Place, 0, len(r.places))
	for _, place := range r.places {
		places = append(places, place)
	}
	return places, nil
}
