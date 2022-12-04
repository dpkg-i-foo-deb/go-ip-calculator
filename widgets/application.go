package widgets

import (
	"ip-calculator/util"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var application *tview.Application

func initApplication() {

	application = tview.NewApplication()
	application.EnableMouse(true)
	application.SetRoot(Pages, true)
	application.SetInputCapture(handleEvent)
}

func handleEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlC:
		Pages.ShowPage("quit-modal")
	default:
		return event
	}
	return nil
}

func StartApplication() {
	err := application.Run()
	util.HandleErrorStop(err)

}

func StopApplication() {
	application.Stop()
	os.Exit(0)
}
