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

const (
	githubUserEventsURL = "https://api.github.com/users/%s/events"
	fieldType           = "type"
)

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

	// Create HTTP client
	client := &http.Client{}

	// Make a new request
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return appError.AppError{
			Err: appError.ErrCallingNewRequest,
		}
	}

	// Add Authorization Bearer header with the Github Personal Access Token as the value to the request
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
		switch eventType := data[fieldType]; eventType {
		case text.CommitCommentEvent:
			helper.CommitCommentEvent(data)
		case text.CreateEvent:
			helper.CreateEvent(data)
		case text.DeleteEvent:
			helper.DeleteEvent(data)
		case text.DiscussionEvent:
			helper.DiscussionEvent(data)
		case text.ForkEvent:
			helper.ForkEvent(data)
		case text.GollumEvent:
			helper.GollumEvent(data)
		case text.IssueCommentEvent:
			helper.IssueCommentEvent(data)
		case text.IssuesEvent:
			helper.IssusesEvent(data)
		case text.MemberEvent:
			helper.MemberEvent(data)
		case text.PublicEvent:
			helper.PublicEvent()
		case text.PullRequestEvent:
			helper.PullRequestEvent(data)
		case text.PullRequestReviewEvent:
			helper.PullRequestReviewEvent(data)
		case text.PullRequestReviewCommentEvent:
			helper.PullRequestReviewCommentEvent(data)
		case text.PushEvent:
			helper.PushEvent(data)
		case text.ReleaseEvent:
			helper.ReleaseEvent(data)
		case text.WatchEvent:
			helper.WatchEvent(data)
		}
	}

	return nil
}
