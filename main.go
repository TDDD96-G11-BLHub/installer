package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("io.github.TDDD96-G11-BLHub.installer")
	w := a.NewWindow("BLHub Installer")

	text := &widget.Label{Text: "Welcome to the BLHub installation wizard."}
	w.SetContent(text)
	w.ShowAndRun()
}
