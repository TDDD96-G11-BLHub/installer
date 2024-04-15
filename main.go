package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/TDDD96-G11-BLHub/installer/internal/assets"
)

func main() {
	a := app.NewWithID("io.github.TDDD96-G11-BLHub.installer")
	w := a.NewWindow("BLHub Installer")

	logo := canvas.NewImageFromResource(assets.BLHubIconDark)
	logo.FillMode = canvas.ImageFillContain

	start := widget.NewButton("Start installation", nil)
	start.Importance = widget.HighImportance
	bottom := container.NewVBox(widget.NewSeparator(), start)

	content := container.NewBorder(nil, bottom, nil, nil, logo)
	card := widget.NewCard("Welcome", "Press the button below to start the installation process.", content)

	w.SetContent(card)
	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
