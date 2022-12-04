package widgets

import (
	"github.com/rivo/tview"
)

var QuitModal *tview.Modal

func initQuitModal() {
	QuitModal = tview.NewModal().SetText("Do you really want to quit?").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(quitModalDoneFunction)
}

func quitModalDoneFunction(buttonindex int, buttonlabel string) {
	if buttonindex == 0 {
		application.Stop()
	} else {
		Pages.SwitchToPage("root")
	}
}
