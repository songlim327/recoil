package gui

import (
	"image/color"
	"recoil/internal/cons"
	"recoil/resources/images"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// subWindow returns a subwindow for general purposes
func subWindow(title string) fyne.Window {
	sw := a.NewWindow(title)
	sw.CenterOnScreen()
	sw.SetIcon(images.Logo512)
	sw.Resize(fyne.NewSize(760, 510))
	sw.SetFixedSize(true)
	return sw
}

// itemList returns a generic list widget
func itemList(data binding.DataList, icon *fyne.StaticResource, entity string) *widget.List {
	return widget.NewListWithData(data,
		// Create item
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(icon), widget.NewLabel(""), layout.NewSpacer(), container.NewGridWithColumns(1, widget.NewButtonWithIcon("", images.Edit, func() {})))
		},
		// Update item
		func(di binding.DataItem, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).Bind(di.(binding.String))
			v, _ := di.(binding.String).Get()

			// edit button
			editBtn := item.(*fyne.Container).Objects[3].(*fyne.Container).Objects[0].(*widget.Button)

			editBtn.OnTapped = func() {
				if entity == cons.BucketEntity {
					editBucketHandler(v)
				} else {
					editKeyHandler(v)
				}
			}
		})
}

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
