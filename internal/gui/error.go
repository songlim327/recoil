package gui

import (
	"recoil/resources/images"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// errorHandler handle error and show it to the user
func errorHandler(err error) {
	const ErrorTitle = "Error"

	ew := a.NewWindow(ErrorTitle)
	ew.CenterOnScreen()
	ew.SetIcon(images.Logo512)
	ew.Resize(fyne.NewSize(200, 140))
	ew.SetFixedSize(true)

	errImg := canvas.NewImageFromResource(images.Error)
	errImg.SetMinSize(fyne.NewSize(64, 64))

	tTitle := centerLabel(ErrorTitle, true)
	tError := centerText(err.Error(), 14, false)

	errorBox := container.NewVBox(tTitle, tError)

	ew.SetContent(container.NewCenter(errorBox))
	ew.Show()
}
