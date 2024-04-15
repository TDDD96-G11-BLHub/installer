package assets

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

var (
	//go:embed blhub-dark.svg
	blhubDark []byte

	//go:embed blhub-light.svg
	blhubLight []byte
)

var BLHubIconDark = &fyne.StaticResource{
	StaticName:    "blhub-dark.svg",
	StaticContent: blhubDark,
}

var BLHubIconLight = &fyne.StaticResource{
	StaticName:    "blhub-light.svg",
	StaticContent: blhubLight,
}
