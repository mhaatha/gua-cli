package web

type PushPayload struct {
	Ref string `json:"ref"`
}

type WatchPayload struct {
	Action string `json:"action"`
}

type CreatePayload struct {
	// Ref can be nil
	Ref *string `json:"ref"`

	RefType      string `json:"ref_type"`
	FullRef      string `json:"full_ref"`
	MasterBranch string `json:"master_branch"`
	Description  string `json:"description"`
	PusherType   string `json:"pusher_type"`
}

type IssuesPayload struct {
	Action string `json:"action"`
}

type PullRequestPayload struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
}

type ForkPayload struct {
	Action string `json:"action"`
	Forkee Forkee `json:"forkee"`
}

type IssueCommentPayload struct {
	Action string `json:"action"`
}

type PullRequestReviewPayload struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
}

type DeletePayload struct {
	Ref        string `json:"ref"`
	RefType    string `json:"ref_type"`
	FullRef    string `json:"full_ref"`
	PusherType string `json:"pusher_type"`
}
