package tui

import (
	"ip-calculator/util"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

)

var application *tview.Application
var title *tview.TextView
var pages *tview.Pages
var quitModal *tview.Modal
var grid *tview.Grid
var ipForm *tview.Form

func init(){

	title = tview.NewTextView()	
	title.SetText("IP Calculator")
	title.SetTextAlign(tview.AlignCenter)
	
	ipForm = tview.NewForm()
	ipForm.SetBorder(true)
	ipForm.SetTitle("Enter the required data").SetTitleAlign(tview.AlignLeft)
	ipForm.AddInputField("Host IP address","",30,nil,nil)
	ipForm.AddInputField("Netmask","",30,nil,nil)
	ipForm.AddButton("Get Values",nil)
	ipForm.AddButton("Clear",nil)

	grid = tview.NewGrid().
		SetColumns(0,0).
		SetRows(2,0,1).
		SetBorders(true).
		AddItem(title,0,0,1,2,0,0,true).		
		AddItem(ipForm,1,0,1,1,0,0,false)

	quitModal = tview.NewModal().SetText("Do you really want to quit?").
		AddButtons([]string {"Yes","No"}).
		SetDoneFunc(quitModalDoneFunction)

	pages = tview.NewPages()

	pages.AddPage("root",grid,true,true)
	pages.AddPage("quit-modal",quitModal,true,false)

	application = tview.NewApplication().EnableMouse(true).SetRoot(pages,true)
	application.SetInputCapture(handleEvent)
	application.SetFocus(ipForm)
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
