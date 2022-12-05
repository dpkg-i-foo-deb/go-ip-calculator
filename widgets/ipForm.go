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
	//Verify the IP address
	addrField := IPForm.GetFormItem(0).(*tview.InputField)

	_, err := ipcalc.ValidateIP(addrField.GetText())

	if err != nil {
		errorModal.SetText(err.Error())
		Pages.ShowPage("error-modal")
		return
	}

	//Verify the net mask

	maskField := IPForm.GetFormItem(1).(*tview.InputField)

	_, err = ipcalc.ValidateNetMask(maskField.GetText())

	if err != nil {
		errorModal.SetText(err.Error())
		Pages.ShowPage("error-modal")
		return
	}

	//Get the results
	//TODO include IPv6 stuff
	results, err := ipcalc.FindIpv4Results(addrField.GetText(), maskField.GetText())

	//Populate the GUI elements using the results 
	setResults(results)
}

func clearIPForm(){
	IPForm.Clear(false)
	IPForm.AddInputField("Host IP address", "", 30, nil, nil)
	IPForm.AddInputField("Netmask", "", 30, nil, nil)
}
