package gui

import (
	"image/color"
	"recoil/resources/images"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

// resBox renders the result container
func resBox() *fyne.Container {
	text := canvas.NewText("Hello world!", color.White)
	return container.NewVBox(text)
}

// keyBox renders the key list container
func keyBox() *fyne.Container {
	text := canvas.NewText("key", color.White)
	return container.NewVBox(text)
}

// opsBox renders the operation container
func opsBox() *fyne.Container {
	l := []fyne.CanvasObject{
		opsButton("Open", images.Add, func() { openDbHandler() }),
		layout.NewSpacer(),
		// TODO: github handler
		opsButton("Github", images.Github, func() {}),
		opsButton("About", images.About, func() { aboutHandler() }),
	}
	return container.NewVBox(l...)
}
