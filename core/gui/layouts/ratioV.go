package layouts

import (
	"fyne.io/fyne/v2"
)

type RatioVLayout struct {
	ratios []float32
}

func NewRatioVLayout(ratios ...float32) *RatioVLayout {
	return &RatioVLayout{ratios: ratios}
}

func (r *RatioVLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	totalRatio := float32(0)

	for _, ratio := range r.ratios {
		totalRatio += ratio
	}

	yOffset := float32(0)
	for i, obj := range objects {
		ratioHeight := (r.ratios[i] / totalRatio) * size.Height
		obj.Resize(fyne.NewSize(size.Width, ratioHeight))
		obj.Move(fyne.NewPos(0, yOffset))
		yOffset += ratioHeight
	}
}

func (r *RatioVLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minWidth := float32(0)
	maxHeight := float32(0)

	for _, obj := range objects {
		min := obj.MinSize()
		if min.Width > minWidth {
			minWidth = min.Width
		}
		maxHeight += min.Height
	}
	return fyne.NewSize(minWidth, maxHeight)
}
