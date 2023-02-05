package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type word struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"word"`
}

func main() {
	response, err := http.Get("www.pantakan.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP Status Code: %d\nBody: %s", response.StatusCode, body)

	if response.StatusCode != 200 {
		os.Exit(1)
	}

	var Words word

	err = json.Unmarshal(body, &Words)
	if err != nil {
		log.Fatal(err)
	}
}
