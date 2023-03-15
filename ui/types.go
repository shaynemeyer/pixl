package ui

import (
	"fyne.io/fyne/v2"

	"github.com/shaynemeyer/pixl/apptype"
	"github.com/shaynemeyer/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}