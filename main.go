package main

import (
	"mrnotifier/configuration"
	"mrnotifier/gitlab"
	"mrnotifier/poll"
	"mrnotifier/tray"
)

func main() {
	readyMRs := make(chan []*gitlab.MergeRequest)

	cfg := configuration.Read()

	gitlabClient := gitlab.GitlabClient{}
	gitlabClient.Init(cfg)

	poll := &poll.Poll{
		MrChangeChan: readyMRs,
		GitlabClient: &gitlabClient,
	}
	go poll.CheckMRStateInLoop()
	tray.Start(readyMRs)
}
