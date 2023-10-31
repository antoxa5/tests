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

	tx, err := connect.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO events (eventID, eventType, userID, eventTime, payload) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	eventID := 1
	eventType := "MyType"
	userID := 100
	eventTime := time.Now()
	payload := "Пробуем кирилицу"

	if _, err := stmt.Exec(eventID, eventType, userID, eventTime, payload); err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Данные успешно добавлены!")
}
