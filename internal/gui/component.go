package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// opsButton returns an extended fyne Button for operation menu
func opsButton(label string, icon fyne.Resource, tapped func()) *widget.Button {
	b := widget.NewButtonWithIcon(label, icon, tapped)
	b.Alignment = widget.ButtonAlignLeading

	return b
}

// centerText returns an extended fyne Text with center position
func centerText(text string, size float32, bold bool) *canvas.Text {
	t := canvas.NewText(text, color.White)
	t.TextStyle.Bold = bold
	t.TextSize = size
	t.Alignment = fyne.TextAlignCenter

	return t
}

// centerLabel returns an extended fyne Label with center position
func centerLabel(text string, bold bool) *widget.Label {
	return widget.NewLabelWithStyle(text, fyne.TextAlignCenter, fyne.TextStyle{Bold: bold})
}
