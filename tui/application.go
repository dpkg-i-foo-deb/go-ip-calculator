package tui

import (
	"ip-calculator/tui/widgets"
	"ip-calculator/util"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var application *tview.Application

var pages *tview.Pages
var quitModal *tview.Modal

func init() {

	quitModal = tview.NewModal().SetText("Do you really want to quit?").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(quitModalDoneFunction)

	pages = tview.NewPages()

	pages.AddPage("root", widgets.MainGrid, true, false)
	pages.AddPage("quit-modal", quitModal, true, false)

	application = tview.NewApplication()
	application.EnableMouse(true)
	application.SetRoot(pages, true)
	application.SetInputCapture(handleEvent)

	pages.SwitchToPage("root")
}

func handleEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlC:
		pages.ShowPage("quit-modal")
	default:
		return event
	}
	return nil
}

func quitModalDoneFunction(buttonindex int, buttonlabel string) {
	if buttonindex == 0 {
		application.Stop()
	} else {
		pages.SwitchToPage("root")
	}
}

func StartApplication() {
	err := application.Run()
	util.HandleErrorStop(err)

}
