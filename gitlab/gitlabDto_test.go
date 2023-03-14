package gitlab

import (
	"testing"
)

func TestParseMergeRequest(t *testing.T) {
	mrs, err := unmarshalMergeRequest([]byte(`[
		{"iid":1,"Upvotes":1,"downvotes":1,"work_in_progress":false,"web_url":"http://x.com/1"},
		{"iid":2,"Upvotes":2,"downvotes":2,"work_in_progress":true,"web_url":"http://x.com/2"}
		]`))

	if err != nil {
		t.Errorf("Error during parsing: %s", err)
		t.FailNow()
	}

	if len(mrs) != 2 {
		t.Errorf("Expected two Merge Requests, but got: %v", mrs)
		t.FailNow()
	}

	mr1 := MergeRequest{Iid: 1, Upvotes: 1, Downvotes: 1, WorkInProgress: false, WebUrl: "http://x.com/1"}
	if *mrs[0] != mr1 {
		t.Errorf("Expected: %v, but got: %v", mr1, mrs[0])
	}

	mr2 := MergeRequest{Iid: 2, Upvotes: 2, Downvotes: 2, WorkInProgress: true, WebUrl: "http://x.com/2"}
	if *mrs[1] != mr2 {
		t.Errorf("Expected: %v, but got: %v", mr2, mrs[1])
	}
}

func TestParseApproval(t *testing.T) {
	approval, err := unmarshalApprovals([]byte(`{"id":1,"project_id":1,
	"approved_by":[
		{"user":{"id":1,"name":"X"}},
		{"user":{"id":2,"name":"Y"}}
	]}`))

	if err != nil {
		t.Errorf("Error during parsing: %s", err)
	}

	if len(approval.ApprovedBy) != 2 {
		t.Errorf("Incorrect number of approvals. Parsed struct: %s", approval)
	}
}
