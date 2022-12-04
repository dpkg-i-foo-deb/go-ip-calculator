package widgets

import "github.com/rivo/tview"

var Title *tview.TextView

func init(){
	Title = tview.NewTextView()	
	Title.SetText("IP Calculator")
	Title.SetTextAlign(tview.AlignCenter)

}
