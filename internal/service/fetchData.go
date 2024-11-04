package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetUsername(username string) {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Getting and using a value from .env
	githubPAT := os.Getenv("PAT")
	username = fmt.Sprintf("https://api.github.com/users/%s/events", username)

	// Make request with the token
	client := &http.Client{}
	req, err := http.NewRequest("GET", username, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+githubPAT)

	// Call the API
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read the responseData
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON response
	var prettyResponse []map[string]interface{}
	err = json.Unmarshal(responseData, &prettyResponse)
	if err != nil {
		log.Fatal(err)
	}

	for _, data := range prettyResponse {
		switch eventType := data["type"]; eventType {
		case "CommitCommentEvent":
			commitCommentEvent(data)
		case "CreateEvent":
			createEvent(data)
		case "DeleteEvent":
			deleteEvent(data)
		case "ForkEvent":
			forkEvent(data)
		case "GollumEvent":
			gollumEvent(data)
		case "IssueCommentEvent":
			issueCommentEvent(data)
		case "IssuesEvent":
			issusesEvent(data)
		case "MemberEvent":
			memberEvent(data)
		case "PublicEvent":
			publicEvent()
		case "PullRequestEvent":
			pullRequestEvent(data)
		case "PullRequestReviewEvent":
			pullRequestReviewEvent(data)
		case "PullRequestReviewCommentEvent":
			pullRequestReviewCommentEvent(data)
		case "PullRequestReviewThreadEvent":
			pullRequestReviewThreadEvent(data)
		case "PushEvent":
			pushEvent(data)
		case "ReleaseEvent":
			releaseEvent(data)
		case "SponsorshipEvent":
			sponsorshipEvent(data)
		case "WatchEvent":
			watchEvent(data)
		}
	}
}
