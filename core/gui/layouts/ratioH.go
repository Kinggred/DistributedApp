package layouts

import "fyne.io/fyne/v2"

type RatioHLayout struct {
	ratios []float32
}

func NewRatioHLayout(ratios ...float32) *RatioHLayout {
	return &RatioHLayout{ratios: ratios}
}

func (r *RatioHLayout) Layout(items []fyne.CanvasObject, size fyne.Size) {
	totalRatio := float32(0)
	for _, ratio := range r.ratios {
		totalRatio += ratio
	}

	xOffset := float32(0)
	for i, item := range items {
		itemWidth := (size.Width * (r.ratios[i] / totalRatio))
		item.Resize(fyne.NewSize(itemWidth, size.Height))
		item.Move(fyne.NewPos(xOffset, 0))
		// CONSIDER: Might want to add some kind of padding
		xOffset += itemWidth
	}
}

func (r *RatioHLayout) MinSize(items []fyne.CanvasObject) fyne.Size {
	minWidth := float32(0)
	maxHeight := float32(0)

	for i, item := range items {
		minWidth += item.MinSize().Width * r.ratios[i]
		height := item.MinSize().Height
		if height > maxHeight {
			maxHeight = height
		}
	}
	return fyne.NewSize(minWidth, maxHeight)
}
