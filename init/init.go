package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const imageDir = "resources\\images"
const genFile = imageDir + "\\image.go"

func main() {
	var files []string

	createResourceFile(genFile)

	err := filepath.Walk(imageDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".png" {
			files = append(files, path)
		}
		return nil
	})
	check(err)

	rf, err := os.OpenFile(genFile, os.O_RDWR|os.O_APPEND, 0660)
	check(err)

	for _, f := range files {
		imgName := filepath.Base(f)
		cImgName := strings.ReplaceAll(imgName, ".png", "")
		caser := cases.Title(language.English)
		varName := strings.ReplaceAll(caser.String(cImgName), "-", "")

		res, err := fyne.LoadResourceFromPath(f)
		check(err)
		rf.WriteString(fmt.Sprintf("%s = %#v\n", varName, res))
	}
	rf.WriteString(")")
	defer rf.Close()
}

// check logs if error
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// createResourceFile creates resource file containing images byte array
func createResourceFile(p string) {
	if fileExist(p) {
		os.Remove(p)
	}

	f, err := os.Create(p)
	check(err)
	f.WriteString("// **** THIS FILE IS AUTO-GENERATED **** //\n\npackage images\n\nimport \"fyne.io/fyne/v2\"\n\nvar (\n")
	defer f.Close()
}

// fileExist check if a file exist
func fileExist(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}
