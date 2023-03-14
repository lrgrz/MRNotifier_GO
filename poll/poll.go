package poll

import (
	"log"
	"mrnotifier/gitlab"
	"sort"
	"time"
)

type Poll struct {
	MrChangeChan chan []*gitlab.MergeRequest
	GitlabClient *gitlab.GitlabClient
}

type mrWithApprovals struct {
	mergeRequest *gitlab.MergeRequest
	approval     *gitlab.Approvals
}

func (p *Poll) CheckMRStateInLoop() {
	prevState := p.getReadyMergeRequests()
	p.MrChangeChan <- prevState
	for {
		time.Sleep(5 * time.Second)
		newState := p.getReadyMergeRequests()
		if stateChanged(prevState, newState) {
			prevState = newState
			p.MrChangeChan <- newState
		}
	}
}

func stateChanged(prev, current []*gitlab.MergeRequest) bool {
	if len(prev) != len(current) {
		return true
	}
	for i, v := range prev {
		if current[i] != v {
			return true
		}
	}
	return false
}

func (p *Poll) getReadyMergeRequests() []*gitlab.MergeRequest {
	mrWithApprovals := p.getMergeRequestsWithApprovalsUntilSuccess()

	var readyMR []*gitlab.MergeRequest

	for _, mrWithApprovals := range mrWithApprovals {
		mr := mrWithApprovals.mergeRequest
		approval := mrWithApprovals.approval

		if mr.Upvotes-mr.Downvotes+len(approval.ApprovedBy) > 0 {
			readyMR = append(readyMR, mr)
		}
	}

	sort.Slice(readyMR, func(i, j int) bool {
		return readyMR[i].Title < readyMR[j].Title
	})

	return readyMR
}

func (p *Poll) getMergeRequestsWithApprovalsUntilSuccess() []*mrWithApprovals {
	for {
		retries := 0
		state, err := p.getMergeRequestsWithApprovals()
		if err == nil {
			if retries > 0 {
				log.Printf("Merge Request fetched after %d retires\n", retries)
			}
			return state
		}
		retries++
		time.Sleep(time.Second)
	}
}

func (p *Poll) getMergeRequestsWithApprovals() ([]*mrWithApprovals, error) {
	var state []*mrWithApprovals

	mergeRequests, err := p.GitlabClient.GetMergeRequest()
	if err != nil {
		log.Printf("Cannot get MergeRequests: %s", err)
		return nil, err
	}

	for _, mr := range mergeRequests {
		approvals, err := p.GitlabClient.GetApprovals(mr)
		if err != nil {
			log.Printf("Cannot get Approvals for MR=%v: %s", mr, err)
			return nil, err
		}

		state = append(state, &mrWithApprovals{mergeRequest: mr, approval: approvals})
	}

	return state, nil
}
