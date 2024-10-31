package layouts

import (
	"math"

	"fyne.io/fyne/v2"
)

type SquareGridLayout struct {
	minSide float32
}

func NewSquareLayout(minSide float32) *SquareGridLayout {
	return &SquareGridLayout{minSide}
}

func (layout *SquareGridLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	side := float32(math.Max(200, math.Min(float64(size.Width), float64(size.Height))))
	squareSize := fyne.NewSize(side/3, side/3) // Divide by 3 for 3x3 grid

	for i, obj := range objects {
		obj.Resize(squareSize)
		xPos := float32(i%3) * squareSize.Width
		yPos := float32(i/3) * squareSize.Height
		obj.Move(fyne.NewPos(xPos, yPos))
	}
}

func (layout *SquareGridLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(layout.minSide, layout.minSide)
}
