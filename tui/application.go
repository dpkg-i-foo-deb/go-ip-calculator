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
var resultsGrid *tview.Grid
var networkAddress *tview.TextView
var broadcastAddress *tview.TextView
var usableRange *tview.TextView
var totalIPAddresses *tview.TextView
var usableIPAddresses *tview.TextView
var ipType *tview.TextView
var clearResultsButton *tview.Form

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
	ipForm.SetItemPadding(2)
	ipForm.SetBorderPadding(1,1,3,3)

	networkAddress = tview.NewTextView()
	broadcastAddress = tview.NewTextView()
	usableRange = tview.NewTextView()
	totalIPAddresses = tview.NewTextView()
	usableIPAddresses = tview.NewTextView()
	ipType = tview.NewTextView()

	clearResultsButton = tview.NewForm().
		AddButton("Clear Results",nil).
		SetButtonsAlign(tview.AlignRight)
	
	clearResultsButton.SetBorderPadding(0,0,0,0)
	clearResultsButton.SetItemPadding(0)

	resultsGrid = tview.NewGrid().
		SetColumns(0,0).
		SetRows(0,0,0,0,0,0,0).
		SetBorders(false).
		SetMinSize(2,2).
		AddItem(tview.NewTextView().SetText("Network Address"),0,0,1,1,1,0,false).
		AddItem(tview.NewTextView().SetText("Broadcast Address"),1,0,1,1,1,0,false).
		AddItem(tview.NewTextView().SetText("Usable Rage"),2,0,1,1,1,0,false).
		AddItem(tview.NewTextView().SetText("Total IP addresses"),3,0,1,1,1,0,false).
		AddItem(tview.NewTextView().SetText("Usable IP addresses"),4,0,1,1,1,0,false).
		AddItem(tview.NewTextView().SetText("IP type"),5,0,1,1,1,0,false).
		AddItem(networkAddress,0,1,1,1,1,0,false).
		AddItem(broadcastAddress,1,1,1,1,1,0,false).
		AddItem(usableRange,2,1,1,1,1,0,false).
		AddItem(totalIPAddresses,3,1,1,1,1,0,false).
		AddItem(usableIPAddresses,4,1,1,1,1,0,false).
		AddItem(ipType,5,1,1,1,1,0,false).
		AddItem(clearResultsButton,6,1,1,1,1,0,false)



		resultsGrid.SetTitle("Results")
		resultsGrid.SetBorder(true)
		resultsGrid.SetBorderPadding(1,1,2,2)
		resultsGrid.SetTitleAlign(tview.AlignCenter)
	

	grid = tview.NewGrid().
		SetColumns(0,0).
		SetRows(2,0,1).
		SetBorders(false).
		AddItem(title,0,0,1,2,0,0,true).		
		AddItem(ipForm,1,0,1,1,0,0,false).
		AddItem(resultsGrid,1,1,1,1,0,0,false)

	grid.SetBorderPadding(0,0,2,2)

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
