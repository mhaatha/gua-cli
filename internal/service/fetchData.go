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
