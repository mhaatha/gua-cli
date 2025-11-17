package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mhaatha/gua-cli/internal/config"
	appError "github.com/mhaatha/gua-cli/internal/errors"
	"github.com/mhaatha/gua-cli/internal/helper"
	"github.com/mhaatha/gua-cli/internal/text"
)

const githubUserEventsURL = "https://api.github.com/users/%s/events"

type FetchDataServiceImpl struct {
	Cfg *config.Config
}

func NewFetchDataService(cfg *config.Config) FetchDataService {
	return &FetchDataServiceImpl{
		Cfg: cfg,
	}
}

func (s *FetchDataServiceImpl) GetUsername(username string) error {
	requestURL := fmt.Sprintf(githubUserEventsURL, username)

	// Make request with the token
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return appError.AppError{
			Err: appError.ErrCallingNewRequest,
		}
	}

	req.Header.Set("Authorization", "Bearer "+s.Cfg.PAT)

	// Call the API
	response, err := client.Do(req)
	if err != nil {
		return appError.AppError{
			Err: appError.ErrNetworkConnectivityProblem,
		}
	}
	if response.StatusCode == 404 {
		return appError.AppError{
			Err: appError.ErrUsernameNotFound,
		}
	}
	if response.StatusCode == 401 {
		return appError.AppError{
			Err: appError.ErrInvalidGithubPAT,
		}
	}
	defer response.Body.Close()

	// Read the responseData
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return appError.AppError{
			Err: appError.ErrReadResponseBody,
		}
	}

	// Unmarshal the JSON response
	var prettyResponse []map[string]interface{}
	err = json.Unmarshal(responseData, &prettyResponse)
	if err != nil {
		return appError.AppError{
			Err: appError.ErrUnmarshalResponseBody,
		}
	}

	for _, data := range prettyResponse {
		switch eventType := data["type"]; eventType {
		case "CommitCommentEvent":
			helper.CommitCommentEvent(data)
		case "CreateEvent":
			helper.CreateEvent(data)
		case "DeleteEvent":
			helper.DeleteEvent(data)
		case "ForkEvent":
			helper.ForkEvent(data)
		case "GollumEvent":
			helper.GollumEvent(data)
		case "IssueCommentEvent":
			helper.IssueCommentEvent(data)
		case "IssuesEvent":
			helper.IssusesEvent(data)
		case "MemberEvent":
			helper.MemberEvent(data)
		case "PublicEvent":
			helper.PublicEvent()
		case "PullRequestEvent":
			helper.PullRequestEvent(data)
		case "PullRequestReviewEvent":
			helper.PullRequestReviewEvent(data)
		case "PullRequestReviewCommentEvent":
			helper.PullRequestReviewCommentEvent(data)
		case "PullRequestReviewThreadEvent":
			helper.PullRequestReviewThreadEvent(data)
		case text.PushEvent:
			helper.PushEvent(data)
		case "ReleaseEvent":
			helper.ReleaseEvent(data)
		case "SponsorshipEvent":
			helper.SponsorshipEvent(data)
		case "WatchEvent":
			helper.WatchEvent(data)
		}
	}

	return nil
}
