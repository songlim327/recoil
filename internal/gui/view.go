package gui

import (
	"recoil/internal/cons"
	"recoil/resources/images"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// keyBox renders the key list container
func keyBox() *fyne.Container {
	keyItemList = itemList(keys, images.Key, cons.KeyEntity)
	eKey := widget.NewEntry()
	eKey.PlaceHolder = "Search for keys..."
	eKey.OnChanged = searchKeyHandler

	return container.NewBorder(eKey, nil, nil, nil, container.NewVScroll(keyItemList))
}

// bucBox renders the bucket list container
func bucBox() *fyne.Container {
	bucketItemList = itemList(buckets, images.Bucket, cons.BucketEntity)
	bucketItemList.OnSelected = bucketHandler
	eBucket := widget.NewEntry()
	eBucket.PlaceHolder = "Search for buckets..."
	eBucket.OnChanged = searchBucketHandler

	return container.NewBorder(eBucket, nil, nil, nil, container.NewVScroll(bucketItemList))
}

// opsBox renders the operation container
func opsBox() *fyne.Container {
	l := append(opsBoxTopView(), layout.NewSpacer())
	l = append(l, opsBoxBottomView()[:]...)

	return container.NewVBox(l...)
}

// dbBox renders the container showing the database name
func dbBox() *fyne.Container {
	return container.NewHBox(
		widget.NewLabelWithData(filename),
	)
}

// opsBoxTopView renders the top container of operation container
func opsBoxTopView() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		opsButton("Open", images.Open, func() {
			openDbHandler()
		}),
		opsButton(cons.Add, images.Add, func() { addHandler() }),
	}
}

// opsBoxBottomView renders the bottom container of operation container
func opsBoxBottomView() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		opsButton("Github", images.Github, func() { githubHandler() }),
		opsButton("About", images.About, func() { aboutHandler() }),
	}
}
