package widgets

import "github.com/rivo/tview"

var MainGrid *tview.Grid

func initMainGrid() {
	MainGrid = tview.NewGrid().
		SetColumns(0, 0).
		SetRows(2, 0, 1).
		SetBorders(false).
		AddItem(Title, 0, 0, 1, 2, 0, 0, true).
		AddItem(IPForm, 1, 0, 1, 1, 0, 0, false).
		AddItem(ResultsGrid, 1, 1, 1, 1, 0, 0, false)

	MainGrid.SetBorderPadding(0, 0, 2, 2)

}
