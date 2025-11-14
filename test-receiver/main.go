package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type UserStatus struct {
	ID             int     `json:"id"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Balance        float64 `json:"balance"`
	ReservedAmount float64 `json:"reserved_amount"`
}

var db *sql.DB

func handleUserStatus(w http.ResponseWriter, r *http.Request) {
	var u UserStatus
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid json", 400)
		return
	}

	_, err := db.Exec(`INSERT INTO user_status (id, latitude, longitude, balance, reserved_amount)
	                   VALUES ($1, $2, $3, $4, $5)`,
		u.ID, u.Latitude, u.Longitude, u.Balance, u.ReservedAmount)
	if err != nil {
		http.Error(w, "db insert failed: "+err.Error(), 500)
		return
	}

	log.Printf("✅ Получены данные от user=%d (резерв: %.2f)", u.ID, u.ReservedAmount)
	w.WriteHeader(http.StatusOK)
}

func exportToCSV() {
	for {
		rows, err := db.Query(`SELECT id, latitude, longitude, balance, reserved_amount, created_at FROM user_status`)
		if err != nil {
			log.Println("Ошибка чтения:", err)
			time.Sleep(30 * time.Second)
			continue
		}
		defer rows.Close()

		file, err := os.Create("/app/analytics.csv")
		if err != nil {
			log.Println("Ошибка создания CSV:", err)
			time.Sleep(30 * time.Second)
			continue
		}
		writer := csv.NewWriter(file)
		writer.Write([]string{"id", "latitude", "longitude", "balance", "reserved_amount", "created_at"})

		for rows.Next() {
			var id int
			var lat, lon, bal, res float64
			var created time.Time
			rows.Scan(&id, &lat, &lon, &bal, &res, &created)
			writer.Write([]string{
				fmt.Sprintf("%d", id),
				fmt.Sprintf("%.6f", lat),
				fmt.Sprintf("%.6f", lon),
				fmt.Sprintf("%.2f", bal),
				fmt.Sprintf("%.2f", res),
				created.Format(time.RFC3339),
			})
		}

		writer.Flush()
		file.Close()
		log.Println("📊 CSV-аналитика обновлена: analytics.csv")
		time.Sleep(30 * time.Second)
	}
}

func main() {
	var err error
	connStr := "postgres://postgres:postgres@db:5432/testdb?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user_status (
			id INT,
			latitude DOUBLE PRECISION,
			longitude DOUBLE PRECISION,
			balance DOUBLE PRECISION,
			reserved_amount DOUBLE PRECISION,
			created_at TIMESTAMP DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatalf("Table create error: %v", err)
	}

	go exportToCSV()

	http.HandleFunc("/api/user/status", handleUserStatus)
	log.Println("🚀 Test Receiver запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
