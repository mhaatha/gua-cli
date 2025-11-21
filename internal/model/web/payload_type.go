package web

type CommitCommentPayload struct {
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

type DeletePayload struct {
	Ref        string `json:"ref"`
	RefType    string `json:"ref_type"`
	FullRef    string `json:"full_ref"`
	PusherType string `json:"pusher_type"`
}
