package tray

import (
	"mrnotifier/gitlab"
	"os/exec"
	"runtime"

	"fyne.io/systray"
)

const MAX_MERGE_REQUESTS = 10

var mrItemsReady = make(chan int)
var mrItems []*systray.MenuItem
var state []*gitlab.MergeRequest

func Start(readyMRs chan []*gitlab.MergeRequest) {
	go monitorState(readyMRs)
	systray.Run(onReady, onExit)
}

func monitorState(readyMRs chan []*gitlab.MergeRequest) {
	<-mrItemsReady
	for {
		state = <-readyMRs
		for i := 0; i < MAX_MERGE_REQUESTS; i++ {
			if i < len(state) {
				mrItems[i].SetTitle(state[i].Title)
				mrItems[i].Show()
			} else {
				mrItems[i].Hide()
			}
		}

		if len(state) > 0 {
			systray.SetIcon(mergeIcon)
		} else {
			systray.SetIcon(gitlabIcon)
		}
	}
}

func onReady() {
	systray.SetIcon(gitlabIcon)
	systray.SetTitle("MR Notifier")
	systray.SetTooltip("Merge request status notifier for GitLab")

	mrItems = make([]*systray.MenuItem, MAX_MERGE_REQUESTS)
	for i := 0; i < MAX_MERGE_REQUESTS; i++ {
		mrItems[i] = systray.AddMenuItem("", "Open Merge request")
		mrItems[i].Hide()
	}

	quit := systray.AddMenuItem("Quit", "Quit the whole app")
	go monitorMenuItems(quit)
	mrItemsReady <- 0
}

func monitorMenuItems(quit *systray.MenuItem) {
	go func() {
		for {
			select {
			case <-mrItems[0].ClickedCh:
				onClicked(0)
			case <-mrItems[1].ClickedCh:
				onClicked(1)
			case <-mrItems[2].ClickedCh:
				onClicked(2)
			case <-mrItems[3].ClickedCh:
				onClicked(3)
			case <-mrItems[4].ClickedCh:
				onClicked(4)
			case <-mrItems[5].ClickedCh:
				onClicked(5)
			case <-mrItems[6].ClickedCh:
				onClicked(6)
			case <-mrItems[7].ClickedCh:
				onClicked(7)
			case <-mrItems[8].ClickedCh:
				onClicked(8)
			case <-mrItems[9].ClickedCh:
				onClicked(9)
			case <-quit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onClicked(i int) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, state[i].WebUrl)
	exec.Command(cmd, args...).Start()
}

func onExit() {
	// clean up here
}
