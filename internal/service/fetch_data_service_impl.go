package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mhaatha/gua-cli/internal/config"
	appError "github.com/mhaatha/gua-cli/internal/errors"
	"github.com/mhaatha/gua-cli/internal/helper"
	"github.com/mhaatha/gua-cli/internal/model/web"
	"github.com/mhaatha/gua-cli/internal/text"
)

const (
	githubUserEventsURL = "https://api.github.com/users/%s/events"
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
	if response.StatusCode == http.StatusNotFound {
		return appError.AppError{
			Err: appError.ErrUsernameNotFound,
		}
	}
	if response.StatusCode == http.StatusUnauthorized {
		return appError.AppError{
			Err: appError.ErrInvalidGithubPAT,
		}
	}
	defer response.Body.Close()

	// Read and decode a streaming array of JSON response data
	dec := json.NewDecoder(response.Body)

	// Read open bracket
	_, err = dec.Token()
	if err != nil {
		return appError.AppError{
			Err: appError.ErrReadNextJSONToken,
		}
	}

	for dec.More() {
		var data web.GithubResponse

		err := dec.Decode(&data)
		if err != nil {
			return appError.AppError{
				Err: appError.ErrDecodeJSON,
			}
		}

		switch data.Type {
		case text.PushEvent:
			helper.PushEvent(data)
		case text.WatchEvent:
			helper.WatchEvent(data)
		case text.CreateEvent:
			helper.CreateEvent(data)
		case text.IssuesEvent:
			helper.IssuesEvent(data)
		case text.PullRequestEvent:
			helper.PullRequestEvent(data)
		case text.ForkEvent:
			helper.ForkEvent(data)
		case text.IssueCommentEvent:
			helper.IssueCommentEvent(data)
		case text.PullRequestReviewEvent:
			helper.PullRequestReviewEvent(data)
		case text.DeleteEvent:
			helper.DeleteEvent(data)
		case text.ReleaseEvent:
			helper.ReleaseEvent(data)
		}
	}

	// Read closing bracket
	_, err = dec.Token()
	if err != nil {
		return appError.AppError{
			Err: appError.ErrReadNextJSONToken,
		}
	}

	return nil
}
