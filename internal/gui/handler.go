package gui

import (
	"fmt"
	"net/url"
	"os"
	"recoil/internal/cons"
	"recoil/internal/core"
	"recoil/resources/images"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pkg/browser"
)

// githubHandler opens repository github link
func githubHandler() {
	err := browser.OpenURL("https://github.com/songlim327/recoil")
	if err != nil {
		errorHandler(err, mw)
	}
}

// settingHandler opens app setting
func settingsHandler() {}

// addHandler opens a new window for creating new bucket/key
func addHandler() {
	f, _ := filename.Get()
	if f == "" {
		dialog.NewInformation(cons.Add, cons.ErrNoDb, mw).Show()
	} else {
		kw := subWindow(cons.Add)

		bsInBytes, _ := buckets.Get()
		bsInstr := []string{}
		for _, v := range bsInBytes {
			bsInstr = append(bsInstr, string(v))
		}

		// Bucket selection when entity equals to key (default hidden)
		lBucket := widget.NewLabel("Bucket")
		sBucket := widget.NewSelect(bsInstr, func(s string) {})
		sBucket.SetSelectedIndex(0)
		sBucket.PlaceHolder = "Select bucket"
		lBucket.Hide()
		sBucket.Hide()

		// Name
		eName := widget.NewEntry()
		eName.Validator = func(text string) error {
			if text == "" {
				return fmt.Errorf("name can't be empty")
			}
			return nil
		}

		// Value (default hidden)
		lValue := widget.NewLabel("Value")
		eValue := widget.NewMultiLineEntry()
		lValue.Hide()
		eValue.Hide()

		// Entity
		sEntity := widget.NewSelect([]string{cons.BucketEntity, cons.KeyEntity}, func(s string) {
			if s == cons.KeyEntity {
				lBucket.Show()
				sBucket.Show()
				lValue.Show()
				eValue.Show()
			} else {
				lBucket.Hide()
				sBucket.Hide()
				lValue.Hide()
				eValue.Hide()
			}
		})
		sEntity.SetSelected(cons.BucketEntity)

		f := widget.NewForm(
			widget.NewFormItem("", container.New(layout.NewFormLayout(),
				widget.NewLabel("Entity"),
				sEntity,
				lBucket,
				sBucket,
				widget.NewLabel("Name"),
				eName,
				lValue,
				eValue,
			)),
		)

		f.SubmitText = "Create"
		f.OnCancel = func() {
			kw.Close()
		}
		f.OnSubmit = func() {
			var err error
			// Validate form
			err = eName.Validate()
			if err != nil {
				errorHandler(err, kw)
				return
			}
			if sEntity.Selected == cons.BucketEntity {
				// Create bucket
				err = addBucket(eName.Text)
			} else {
				// Create key
				err = updateKey(sBucket.Selected, eName.Text, eValue.Text)
			}

			if err != nil {
				errorHandler(err, kw)
			} else {
				dialog.NewInformation(cons.Add, cons.AddMsg, mw).Show()
				// Refresh item list after add
				if sEntity.Selected == cons.BucketEntity {
					err = bindAllBuckets()
				} else {
					err = bindAllKeys(selBucket)
				}
				// Error handler after refresh buckets/keys
				if err != nil {
					errorHandler(err, mw)
				}

				kw.Close()
			}
		}

		kw.SetContent(f)
		kw.Show()
	}
}

// deleteBucketHandler handles the delete of a bucket
func deleteBucketHandler(item string) {
	f, _ := filename.Get()
	if f == "" {
		dialog.NewInformation(cons.BucketDelete, cons.ErrNoDb, mw).Show()
	} else if item == "" {
		dialog.NewInformation(cons.BucketDelete, cons.ErrNoBucket, mw).Show()
	} else {
		dialog.NewConfirm(cons.BucketDelete, fmt.Sprintf("Delete %v?", item), func(b bool) {
			var err error
			if b {
				err = deleteBucket(item)

				// Delete bucket error handling
				if err != nil {
					errorHandler(err, mw)
				} else {
					err := bindAllBuckets()
					if err != nil {
						errorHandler(err, mw)
					}
					dialog.NewInformation(cons.BucketDelete, cons.DeleteMsg, mw).Show()
				}
			}
		}, mw).Show()
	}
}

// deleteKeyHandler handles the delete of a key
func deleteKeyHandler(item string) {
	f, _ := filename.Get()
	if f == "" {
		dialog.NewInformation(cons.KeyDelete, cons.ErrNoDb, mw).Show()
	} else if item == "" {
		dialog.NewInformation(cons.KeyDelete, cons.ErrNoKey, mw).Show()
	} else {
		dialog.NewConfirm(cons.KeyDelete, fmt.Sprintf("Delete %v?", item), func(b bool) {
			var err error
			if b {
				err = deleteKey(selBucket, item)

				// Delete key error handling
				if err != nil {
					errorHandler(err, mw)
				} else {
					err := bindAllKeys(selBucket)
					if err != nil {
						errorHandler(err, mw)
					}
					dialog.NewInformation(cons.KeyDelete, cons.DeleteMsg, mw).Show()
				}
			}
		}, mw).Show()
	}
}

// editBucketHandler open a new window for editing bucket name
func editBucketHandler() {
	f, _ := filename.Get()
	if f == "" {
		dialog.NewInformation(cons.BucketEdit, cons.ErrNoDb, mw).Show()
	} else if selBucket == "" {
		dialog.NewInformation(cons.BucketEdit, cons.ErrNoBucket, mw).Show()
	} else {
		bw := subWindow(selBucket)
		lBucket := widget.NewLabel(selBucket)
		eBucket := widget.NewEntry()
		eBucket.SetPlaceHolder("Enter a new bucket name")
		eBucket.Validator = func(text string) error {
			if text == "" {
				return fmt.Errorf("new name can't be empty")
			}
			return nil
		}

		f := widget.NewForm(
			widget.NewFormItem("Name", lBucket),
			widget.NewFormItem("New Name", eBucket),
		)
		f.SubmitText = "Save"
		f.OnCancel = func() {
			bw.Close()
		}
		f.OnSubmit = func() {
			err := f.Validate()
			if err != nil {
				errorHandler(err, bw)
			}

			err = updateBucket(selBucket, eBucket.Text)
			if err != nil {
				errorHandler(err, bw)
			} else {
				err := bindAllBuckets()
				if err != nil {
					errorHandler(err, mw)
				}
				dialog.NewInformation(cons.BucketEdit, cons.BucketEditMsg, mw).Show()
				bw.Close()
			}
		}

		bw.SetContent(f)
		bw.Show()
	}
}

// editKeyHandler open a new window and display key value
func editKeyHandler() {
	f, _ := filename.Get()
	if f == "" {
		dialog.NewInformation(cons.KeyEdit, cons.ErrNoDb, mw).Show()
	} else if selKey == "" {
		dialog.NewInformation(cons.KeyEdit, cons.ErrNoKey, mw).Show()
	} else {
		// Define textarea widget
		textArea := widget.NewMultiLineEntry()
		textArea.Wrapping = fyne.TextWrapWord
		textArea.SetMinRowsVisible(20)

		// Get key value and set to text area
		v, err := db.GetKey(selBucket, selKey)
		if err != nil {
			errorHandler(err, mw)
		}
		textArea.SetText(string(v))

		kw := subWindow(selKey)
		f := widget.NewForm(widget.NewFormItem("Value", textArea))
		f.SubmitText = "Save"
		f.OnCancel = func() {
			kw.Close()
		}
		f.OnSubmit = func() {
			err := updateKey(selBucket, selKey, textArea.Text)
			if err != nil {
				errorHandler(err, kw)
			} else {
				dialog.NewInformation(cons.KeyEdit, cons.KeyEditMsg, mw).Show()
				kw.Close()
			}
		}

		kw.SetContent(f)
		kw.Show()
	}
}

// keyHandler retrieve key and set to global variable
func keyHandler(id widget.ListItemID) {
	v, _ := keys.GetValue(id)
	selKey = string(v)
}

// bucketHandler opens bolt database bucket
func bucketHandler(id widget.ListItemID) {
	// Get selected bucket
	v, err := buckets.GetValue(id)
	if err != nil {
		errorHandler(err, mw)
	}
	// Refresh bucket keys
	err = bindAllKeys(v)
	if err != nil {
		errorHandler(err, mw)
	}

	// Clear key item list selected value
	selKey = ""
	keyItemList.UnselectAll()
}

// openDbHandler opens a new bolt database connection to existing/new database
func openDbHandler() {
	d := dialog.NewFileOpen(func(uc fyne.URIReadCloser, callbackErr error) {
		if callbackErr != nil {
			errorHandler(callbackErr, mw)
			return
		}

		// Proceed to create/open if a file is selected
		if uc != nil {
			f := uc.URI().Path()

			var err error
			db, err = core.New(f)
			if err != nil {
				errorHandler(err, mw)
			}

			// Bind filename string
			filename.Set(uc.URI().Name())
			err = bindAllBuckets()
			if err != nil {
				errorHandler(err, mw)
			}
		}
	}, mw)

	curDir, err := os.Getwd()
	if err != nil {
		errorHandler(err, mw)
	}

	// Current app directory
	fileUri := storage.NewFileURI(curDir)
	fileUriLister, err := storage.ListerForURI(fileUri)
	if err != nil {
		errorHandler(err, mw)
	}

	// Init a .db file type for URI
	fileExt := storage.NewExtensionFileFilter([]string{".db"})

	d.SetLocation(fileUriLister)
	d.SetFilter(fileExt)
	d.Resize(fyne.NewSize(760, 540))
	d.Show()
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

// errorHandler handle error and show it to the user
func errorHandler(err error, parentWindow fyne.Window) {
	ew := dialog.NewError(err, parentWindow)
	ew.Resize(fyne.NewSize(200, 140))

	ew.Show()
}
