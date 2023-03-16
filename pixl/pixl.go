package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/shaynemeyer/pixl/apptype"
	"github.com/shaynemeyer/pixl/pxcanvas"
	"github.com/shaynemeyer/pixl/swatch"
	"github.com/shaynemeyer/pixl/ui"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := apptype.State {
		BrushColor: color.NRGBA{255,255,255,255},
		SwatchSelected: 0,
	}

	pixlCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea: fyne.NewSize(600,600),
		CanvasOffset: fyne.NewPos(0,0),
		PxRows: 10,
		PxCols: 10,
		PxSize: 30,
	}

	pixlCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)


	appInit := ui.AppInit{
		PixlCanvas: pixlCanvas,
		PixlWindow: pixlWindow,
		State: &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	
	appInit.PixlWindow.ShowAndRun()
}