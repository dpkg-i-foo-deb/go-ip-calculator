package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var errorModal *tview.Modal

func initErrorModal(){
	errorModal = tview.NewModal().
					SetText("Some error").
					AddButtons([]string{"Ok"}).
					SetDoneFunc(errorModalDoneFunction).
					SetBackgroundColor(tcell.ColorRed)
}

func errorModalDoneFunction (buttonIndex int, buttonLabel string){
	if buttonIndex == 0 {
		Pages.SwitchToPage("root")
	}
}
