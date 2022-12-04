package widgets

import "github.com/rivo/tview"

var ResultsGrid *tview.Grid
var NetworkAddress *tview.TextView
var BroadcastAddress *tview.TextView
var UsableRange *tview.TextView
var TotalIPAddresses *tview.TextView
var UsableIPAddresses *tview.TextView
var IPType *tview.TextView

func initResults() {
	NetworkAddress = tview.NewTextView()
	BroadcastAddress = tview.NewTextView()
	UsableRange = tview.NewTextView()
	TotalIPAddresses = tview.NewTextView()
	UsableIPAddresses = tview.NewTextView()
	IPType = tview.NewTextView()

	ResultsGrid = tview.NewGrid().
		SetColumns(0, 0).
		SetRows(0, 0, 0, 0, 0, 0, 0).
		SetBorders(false).
		SetMinSize(2, 2).
		AddItem(tview.NewTextView().SetText("Network Address"), 0, 0, 1, 1, 1, 0, false).
		AddItem(tview.NewTextView().SetText("Broadcast Address"), 1, 0, 1, 1, 1, 0, false).
		AddItem(tview.NewTextView().SetText("Usable Rage"), 2, 0, 1, 1, 1, 0, false).
		AddItem(tview.NewTextView().SetText("Total IP addresses"), 3, 0, 1, 1, 1, 0, false).
		AddItem(tview.NewTextView().SetText("Usable IP addresses"), 4, 0, 1, 1, 1, 0, false).
		AddItem(tview.NewTextView().SetText("IP type"), 5, 0, 1, 1, 1, 0, false).
		AddItem(NetworkAddress, 0, 1, 1, 1, 1, 0, false).
		AddItem(BroadcastAddress, 1, 1, 1, 1, 1, 0, false).
		AddItem(UsableRange, 2, 1, 1, 1, 1, 0, false).
		AddItem(TotalIPAddresses, 3, 1, 1, 1, 1, 0, false).
		AddItem(UsableIPAddresses, 4, 1, 1, 1, 1, 0, false).
		AddItem(IPType, 5, 1, 1, 1, 1, 0, false).
		AddItem(ClearResultsButton, 6, 1, 1, 1, 1, 0, false)

	ResultsGrid.SetTitle("Results")
	ResultsGrid.SetBorder(true)
	ResultsGrid.SetBorderPadding(1, 1, 2, 2)
	ResultsGrid.SetTitleAlign(tview.AlignCenter)
}
