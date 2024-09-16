package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetUsername(username string) {
	username = fmt.Sprintf("https://api.github.com/users/%s/events", username)

	// Call the API
	response, err := http.Get(username)
	if err != nil {
		log.Fatal(err)
	}

	// Read the responseData
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Pretty-print the JSON response
	var prettyResponse []map[string]interface{}
	err = json.Unmarshal(responseData, &prettyResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prettyResponse[0]["type"])
}
