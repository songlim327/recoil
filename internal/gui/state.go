package gui

import (
	"recoil/internal/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Global variables for ui state
var (
	a           fyne.App
	mw          fyne.Window
	db          *core.Database
	selBucket   string
	selKey      string
	filename    binding.String
	buckets     binding.StringList
	keys        binding.StringList
	keyItemList *widget.List
)
