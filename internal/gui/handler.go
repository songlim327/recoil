package gui

import (
	"fmt"
	"net/url"
	"recoil/internal/cons"
	"recoil/resources/images"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// openDbHandler opens a new bolt database connection to existing/new database
func openDbHandler() {
	dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		fmt.Println(uc.URI().Name())

		if err != nil {
			errorHandler(err)
		}
	}, mw).Show()
}

// aboutHandler open a new small window showing information about the app
func aboutHandler() {
	aw := a.NewWindow("About")
	aw.CenterOnScreen()
	aw.SetIcon(images.Logo512)
	aw.Resize(fyne.NewSize(360, 270))
	aw.SetFixedSize(true)

	logo := canvas.NewImageFromResource(images.Logo512)
	logo.SetMinSize(fyne.NewSize(64, 64))

	tName := centerLabel(cons.AppName, true)
	tDesc := centerText(cons.AppDesc, 14, false)
	tCopyright := centerText(fmt.Sprintf("Copyright © %v %v", cons.Year, cons.Author), 12, false)

	u, _ := url.Parse("mailto:" + cons.Author)
	hAuthor := widget.NewHyperlink(cons.Author, u)
	hAuthor.Alignment = fyne.TextAlignCenter

	// Attribution
	tAttr := centerText("This app used icons created by Freepik - Flaticon", 12, false)

	aboutBox := container.NewVBox(
		container.NewCenter(logo),
		tName,
		tDesc,
		tCopyright,
		hAuthor,
		tAttr,
	)
	aboutCard := widget.NewCard("", "", aboutBox)

	aw.SetContent(container.NewCenter(aboutCard))
	aw.Show()
}
