package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
)

func main() {
	dsn := "http://127.0.0.1:8123/default"
	connect, err := sql.Open("clickhouse", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = connect.Ping(); err != nil {
		log.Fatal(err)
	}

	events, err := fetchEventsByTypeAndDateRange(connect, "opa", "2022-04-01", "2023-04-30")
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		fmt.Printf("ID: %d, Type: %s, UserID: %d, Time: %s, Payload: %s\n", event.eventID, event.eventType, event.userID, event.eventTime, event.payload)
	}
}

type Event struct {
	eventID   int
	eventType string
	userID    int
	eventTime time.Time
	payload   string
}

func fetchEventsByTypeAndDateRange(db *sql.DB, eventType, startDate, endDate string) ([]Event, error) {
	query := `
		SELECT eventID, eventType, userID, eventTime, payload 
		FROM events 
		WHERE eventType = ? AND eventTime BETWEEN ? AND ?
	`
	rows, err := db.Query(query, eventType, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.eventID, &event.eventType, &event.userID, &event.eventTime, &event.payload); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, rows.Err()
}
