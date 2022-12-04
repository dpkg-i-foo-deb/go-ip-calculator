package widgets

import "github.com/rivo/tview"

var Pages *tview.Pages

func initPages() {
	Pages = tview.NewPages()
	Pages.AddPage("root", MainGrid, true, true)
	Pages.AddPage("quit-modal", QuitModal, true, false)
	Pages.AddPage("error-modal",errorModal,true,false)
}
