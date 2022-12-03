package tui

import (
	"ip-calculator/util"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

)

var application *tview.Application
var box *tview.Box
var quitModal *tview.Modal

func init(){
	box = tview.NewBox().SetBorder(true).SetTitle("IP Calculator")
	quitModal = tview.NewModal().SetText("Do you really want to quit?").
		AddButtons([]string {"Yes","No"}).
		SetDoneFunc(quitModalDoneFunction)	
	application = tview.NewApplication().EnableMouse(true).SetRoot(box,true)
	application.SetInputCapture(handleEvent)
}

func handleEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key(){
		case tcell.KeyCtrlC:
			application.SetRoot(quitModal,false)
		default:
			return event
	}
	return nil
}

func quitModalDoneFunction(buttonindex int, buttonlabel string){
	if buttonindex == 0{
		application.Stop()
	}else{
		application.SetRoot(box,true)
	}
}

func StartApplication(){
	err := application.Run()
	util.HandleErrorStop(err)

}
