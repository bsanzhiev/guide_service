package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bsanzhiev/guide_service/domain"
	"github.com/bsanzhiev/guide_service/ports"
)

type PlaceHandler struct {
	Service ports.PlaceService
}

// NewPlaceHandler создает новый обработчик для мест
func NewPlaceHandler(service ports.PlaceService) *PlaceHandler {
	return &PlaceHandler{
		Service: service,
	}
}

// RegisterRoutes регистрирует все маршруты для сервера
func (h *PlaceHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/places", h.handleGetAllPlaces)
	mux.HandleFunc("/places/", h.handleGetPlaceByID)
}

// handleGetAllPlaces обрабатывает запрос на получение всех мест
func (h *PlaceHandler) handleGetAllPlaces(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	places, err := h.Service.GetAllPlaces()
	if err != nil {
		http.Error(w, "Error retrieving places", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(places)
}

// handlePlaceById обрабатывает запросы на получение, обновление и удаление места по ID
func (h *PlaceHandler) handleGetPlaceByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/places/"):]

	switch r.Method {
	case http.MethodGet:
		place, err := h.Service.GetPlaceByID(id)
		if err != nil {
			http.Error(w, "Place not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(place)

	case http.MethodPut:
		var place domain.Place
		err := json.NewDecoder(r.Body).Decode(&place)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		updatedPlace, err := h.Service.UpdatePlace(id, place)
		if err != nil {
			http.Error(w, "Error updating place", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedPlace)

	case http.MethodDelete:
		err := h.Service.DeletePlace(id)
		if err != nil {
			http.Error(w, "Error deleting place", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK) // or StatusNoContent

	default:
		http.Error(w, "Metod not allowed", http.StatusMethodNotAllowed)
	}
}

func StartServer(handler *PlaceHandler, port int) error {
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	server := http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: mux,
	}
	return server.ListenAndServe()
}
