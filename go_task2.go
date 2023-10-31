package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
)

func main() {
	connect, err := sql.Open("clickhouse", "tcp://127.0.0.1:8123?debug=true")
	if err != nil {
		log.Fatal(err)
	}

	if err := connect.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = connect.Exec(`
		INSERT INTO events (eventID, eventType, userID, eventTime, payload)
		VALUES (?, ?, ?, ?, ?)`,
		1, "login", 100, time.Now(), "Some payload data for login",
	)

	if err != nil {
		log.Fatal(err)
	}
}
