package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas *PxCanvas
	canvasImage *canvas.Image
	canvasBorder []canvas.Line
}

// WidgetRenderer interface implementation.
func (renderer *PxCanvasRenderer) MinSize() fyne.Size {
	return renderer.pxCanvas.DrawingArea
}

// WidgetRenderer interface implementation.
func (renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)
	for i := 0; i < len(renderer.canvasBorder); i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}
	objects = append(objects, renderer.canvasImage)
	return objects
}

// WidgetRenderer interface implementation.
func (renderer *PxCanvasRenderer) Destroy() {}

// WidgetRenderer interface implementation.
func (renderer *PxCanvasRenderer) Layout(size fyne.Size) {}

func (renderer *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := renderer.pxCanvas.PxCols
	imgPxHeight := renderer.pxCanvas.PxRows
	pxSize := renderer.pxCanvas.PxSize
	renderer.canvasImage.Move(fyne.NewPos(renderer.pxCanvas.CanvasOffset.X, renderer.pxCanvas.CanvasOffset.Y))
	renderer.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeight*pxSize)))
}

func (renderer *PxCanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := renderer.pxCanvas.CanvasOffset
	imgWidth := renderer.canvasImage.Size().Width
	imgHeight := renderer.canvasImage.Size().Height

	left := &renderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	top := &renderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X + imgWidth, offset.Y)
	
	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X + imgWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X + imgWidth, offset.Y+imgHeight)
	
	bottom := &renderer.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y + imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X + imgWidth, offset.Y + imgHeight)
}