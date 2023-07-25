package gui

import (
	"recoil/internal/cons"
	"recoil/resources/images"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

// Initialize package gui state
func init() {
	filename = binding.NewString()
	buckets = binding.NewStringList()
	keys = binding.NewStringList()
}

// CreateApp create the base window for the application
func CreateApp() {
	a = app.New()
	a.Settings().SetTheme(&CustomTheme{})

	mw = a.NewWindow(cons.AppName)
	mw.CenterOnScreen()
	mw.SetMaster()
	mw.SetIcon(images.Logo512)
	mw.Resize(fyne.NewSize(800, 550))
	mw.SetContent(mainLayout())
	mw.ShowAndRun()
}

// mainLayout returns the main layout of the application
func mainLayout() *fyne.Container {
	keyLayout := container.NewHSplit(
		bucBox(),
		keyBox(),
	)
	// split layout in 1:1
	keyLayout.SetOffset(0.5)

	return container.NewBorder(
		dbBox(), nil, opsBox(), nil,
		keyLayout,
	)
}
