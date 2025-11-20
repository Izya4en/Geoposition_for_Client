package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Драйвер Postgres

	"geocash/internal/analytics"
	"geocash/internal/dashboard"
	"geocash/internal/platform/postgres"
)

func main() {
	// 1. Подключение к БД (Конфигурация должна браться из config.yaml)
	connStr := "user=postgres password=secret dbname=atm_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка соединения
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// 2. Инициализация слоев (Dependency Injection)

	// Слой данных (Repository)
	analyticsRepo := postgres.NewAnalyticsRepository(db)

	// Слой бизнес-логики (Service)
	// Сервис получает на вход репозиторий
	analyticsService := analytics.NewService(analyticsRepo)

	// Слой представления (Handler)
	// Хэндлер получает на вход сервис
	dashHandler := dashboard.NewHandler(analyticsService)

	// 3. Роутинг
	http.HandleFunc("/api/v1/efficiency", dashHandler.GetTerminalEfficiency)

	// 4. Запуск сервера
	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
