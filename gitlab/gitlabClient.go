package gitlab

import (
	"fmt"
	"log"
	"mrnotifier/configuration"
	"net/url"
)

type GitlabClient struct {
	baseUrl  *url.URL
	httpUtil *httpUtil
}

func (g *GitlabClient) Init(conf *configuration.Config) {
	var err error
	g.baseUrl, err = url.Parse(conf.BaseUrl)
	if err != nil {
		log.Fatalf("Incorrect Gitlab URL: %s", conf.BaseUrl)
	}

	g.httpUtil = &httpUtil{}
	g.httpUtil.configure(conf.VerifySSL, conf.ApiToken)
}

func (g *GitlabClient) GetMergeRequest() ([]*MergeRequest, error) {
	bytes, err := g.httpUtil.get(g.buildMergeRequestQuery())
	if err != nil {
		return nil, err
	}
	return unmarshalMergeRequest(bytes)
}

func (g *GitlabClient) buildMergeRequestQuery() string {
	url := g.baseUrl.JoinPath("/api/v4/merge_requests")

	query := url.Query()
	query.Add("state", "opened")
	query.Add("scope", "assigned_to_me")
	url.RawQuery = query.Encode()

	return url.String()
}

func (g *GitlabClient) GetApprovals(mr *MergeRequest) (*Approvals, error) {
	bytes, err := g.httpUtil.get(g.buildApprovalsQuery(mr))
	if err != nil {
		return nil, err
	}
	return unmarshalApprovals(bytes)
}

func (g *GitlabClient) buildApprovalsQuery(mr *MergeRequest) string {
	path := fmt.Sprintf("/api/v4/projects/%d/merge_requests/%d/approvals", mr.ProjectId, mr.Iid)
	url := g.baseUrl.JoinPath(path)
	return url.String()
}
