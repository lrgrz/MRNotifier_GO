package gitlab

import "encoding/json"

type MergeRequest struct {
	Iid            int    `json:"iid"`
	ProjectId      int    `json:"project_id"`
	Title          string `json:"title"`
	WorkInProgress bool   `json:"work_in_progress"`
	Upvotes        int    `json:"upvotes"`
	Downvotes      int    `json:"downvotes"`
	WebUrl         string `json:"web_url"`
}

type Approvals struct {
	ApprovedBy []ApprovedBy `json:"approved_by"`
}

type ApprovedBy struct {
}

func unmarshalMergeRequest(jsonBody []byte) ([]*MergeRequest, error) {
	var mrs []*MergeRequest
	err := json.Unmarshal(jsonBody, &mrs)
	return mrs, err
}

func unmarshalApprovals(jsonBody []byte) (*Approvals, error) {
	approval := Approvals{}
	err := json.Unmarshal(jsonBody, &approval)
	return &approval, err
}
