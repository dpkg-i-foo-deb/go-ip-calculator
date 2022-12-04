package widgets

import "github.com/rivo/tview"

var ClearResultsButton *tview.Form

func init(){
	ClearResultsButton = tview.NewForm().
		AddButton("Clear Results",nil).
		SetButtonsAlign(tview.AlignRight)
	
	ClearResultsButton.SetBorderPadding(0,0,0,0)
	ClearResultsButton.SetItemPadding(0)

}
