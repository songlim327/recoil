package gui

import (
	"recoil/internal/cons"
	"recoil/resources/images"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func CreateApp() {
	a = app.New()
	mw = a.NewWindow(cons.AppName)

	mw.CenterOnScreen()
	mw.SetMaster()
	mw.SetIcon(images.Logo512)
	mw.Resize(fyne.NewSize(800, 550))

	mw.SetContent(container.NewHBox(
		opsBox(),
		keyBox(),
		resBox(),
	))

	mw.ShowAndRun()
}
