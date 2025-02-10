package main

import (
	"log"
	"net/http"
	"os"

	http_server "github.com/bsanzhiev/guide_service/adapters/http"
	"github.com/bsanzhiev/guide_service/adapters/memory"
	"github.com/bsanzhiev/guide_service/domain"
	"github.com/bsanzhiev/guide_service/ports"
)

func main() {
	// Инициализация репозитория
	repo := memory.NewPlaceRepositoryMemory()

	// Создание сервиса, использующего репозиторий
	service := &domain.PlaceService{
		Repository: repo,
	}

	// Создание и настройка HTTP обработчика
	handler := http_server.NewPlaceHandler(service)

	// Определение порта для сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Используем порт по умолчанию, если переменная окружения не задана
	}

	// Запуск HTTP сервера
	log.Printf("Starting server on port %s", port)
	err := http_server.StartServer(handler, port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
