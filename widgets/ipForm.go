package widgets

import (
	"ip-calculator/ipcalc"
	"github.com/rivo/tview"
)

var IPForm *tview.Form

func initIPForm() {
	IPForm = tview.NewForm()
	IPForm.SetBorder(true)
	IPForm.SetTitle("Enter the required data").SetTitleAlign(tview.AlignLeft)
	IPForm.AddInputField("Host IP address", "", 30, nil, nil)
	IPForm.AddInputField("Netmask", "", 30, nil, nil)
	IPForm.AddButton("Get Values", getValues)
	IPForm.AddButton("Clear", clearIPForm)
	IPForm.SetItemPadding(2)
	IPForm.SetBorderPadding(1, 1, 3, 3)

}

func getValues(){
	//Step 1, verify the IP address
	addrField := IPForm.GetFormItem(0).(*tview.InputField)

	_, err := ipcalc.ValidateIP(addrField.GetText())

	if err != nil {
		errorModal.SetText(err.Error())
		Pages.ShowPage("error-modal")
	}

	//TODO wait for the user to press OK before clearing

	//clearIPForm()
}

func clearIPForm(){
	IPForm.Clear(false)
	IPForm.AddInputField("Host IP address", "", 30, nil, nil)
	IPForm.AddInputField("Netmask", "", 30, nil, nil)
}
