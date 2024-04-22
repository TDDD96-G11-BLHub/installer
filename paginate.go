package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*PageSelector)(nil)

type PageSelector struct {
	widget.BaseWidget
}

func (p *PageSelector) CreateRenderer() fyne.WidgetRenderer {
	return nil
}
