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
var grid *tview.Grid

func init(){
	grid = tview.NewGrid().
		SetColumns(0,0).
		SetRows(2,0,1).
		SetBorders(false).
		AddItem(widgets.Title,0,0,1,2,0,0,true).		
		AddItem(widgets.IPForm,1,0,1,1,0,0,false).
		AddItem(widgets.ResultsGrid,1,1,1,1,0,0,false)

	grid.SetBorderPadding(0,0,2,2)

	quitModal = tview.NewModal().SetText("Do you really want to quit?").
		AddButtons([]string {"Yes","No"}).
		SetDoneFunc(quitModalDoneFunction)

	pages = tview.NewPages()

	pages.AddPage("root",grid,true,true)
	pages.AddPage("quit-modal",quitModal,true,false)

	application = tview.NewApplication().EnableMouse(true).SetRoot(pages,true)
	application.SetInputCapture(handleEvent)
	application.SetFocus(widgets.IPForm)
}

func handleEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key(){
		case tcell.KeyCtrlC:
			pages.ShowPage("quit-modal")
		default:
			return event
	}
	return nil
}

func quitModalDoneFunction(buttonindex int, buttonlabel string){
	if buttonindex == 0{
		application.Stop()
	}else{
		pages.SwitchToPage("root")
	}
}

func StartApplication(){
	err := application.Run()
	util.HandleErrorStop(err)

}
