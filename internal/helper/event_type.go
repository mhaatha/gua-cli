package helper

import (
	"encoding/json"

	"github.com/mhaatha/gua-cli/internal/model/web"
)

func PushEvent(rawData web.GithubResponse) {
	var payload web.PushPayload

	json.Unmarshal(rawData.Payload, &payload)
}

func WatchEvent(rawData web.GithubResponse) {
	var payload web.WatchPayload

	json.Unmarshal(rawData.Payload, &payload)
}

func CreateEvent(rawData web.GithubResponse) {
	var payload web.CreatePayload

	json.Unmarshal(rawData.Payload, &payload)
}

func IssuesEvent(rawData web.GithubResponse) {
	var payload web.IssuesPayload

	json.Unmarshal(rawData.Payload, &payload)
}

func PullRequestEvent(rawData web.GithubResponse) {
	var payload web.PullRequestPayload

	json.Unmarshal(rawData.Payload, &payload)
}

func ForkEvent(rawData web.GithubResponse) {
	var payload web.ForkPayload

	json.Unmarshal(rawData.Payload, &payload)
}

func IssueCommentEvent(rawData web.GithubResponse) {
	var payload web.IssueCommentPayload

	json.Unmarshal(rawData.Payload, &payload)
}

func PullRequestReviewEvent(rawData web.GithubResponse) {
	var payload web.PullRequestReviewPayload

	json.Unmarshal(rawData.Payload, &payload)
}

func DeleteEvent(rawData web.GithubResponse) {
	var payload web.DeletePayload

	json.Unmarshal(rawData.Payload, &payload)
}

func ReleaseEvent(rawData web.GithubResponse) {
	var payload web.ReleasePayload

	json.Unmarshal(rawData.Payload, &payload)
}
