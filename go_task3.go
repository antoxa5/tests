package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
)

type Event struct {
	EventType string `json:"eventType"`
	UserID    int    `json:"userID"`
	EventTime string `json:"eventTime"`
	Payload   string `json:"payload"`
}

func main() {
	http.HandleFunc("/api/event", handleEvent)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = insertEventToClickHouse(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertEventToClickHouse(event Event) error {
	dsn := "http://127.0.0.1:8123/default"
	connect, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return err
	}

	if err = connect.Ping(); err != nil {
		return err
	}

	eventTime, err := time.Parse("2006-01-02 15:04:05", event.EventTime)
	if err != nil {
		return err
	}

	_, err = connect.Exec(
		`INSERT INTO events (eventType, userID, eventTime, payload) VALUES (?, ?, ?, ?)`,
		event.EventType, event.UserID, eventTime, event.Payload,
	)
	if err != nil {
		return err
	}

	return nil
}
