package ui

import (
	"fyne.io/fyne/v2"

	"github.com/shaynemeyer/pixl/apptype"
	"github.com/shaynemeyer/pixl/pxcanvas"
	"github.com/shaynemeyer/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}