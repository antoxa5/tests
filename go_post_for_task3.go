package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:8080/api/event"

	event := `
	{
	  "eventType": "newType",
	  "userID": 1,
	  "eventTime": "2023-04-09 13:00:00",
	  "payload": "Новый PayLoad"
	}
	`

	response, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(event)))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return
	}

	data, _ := io.ReadAll(response.Body)
	fmt.Println("Response:", string(data))
}
