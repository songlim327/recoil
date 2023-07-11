package main

import (
	"fmt"
	"log"
	"recoil/internal/core"
	"recoil/internal/gui"
)

var d *core.Database

func main() {
	d, err := core.New("./qchef.db")
	if err != nil {
		log.Fatal(err)
	}

	b, err := d.IterateBucket()
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range b {
		fmt.Println(string(bucket))
	}

	gui.CreateApp()
}
