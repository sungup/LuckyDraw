package main

import "fyne.io/fyne/v2"

type LuckyDrawLayout struct {
	title fyne.CanvasObject
	body  fyne.CanvasObject
}

func (l *LuckyDrawLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	titleMinSz := l.title.MinSize()
	bodyMinSz := l.body.MinSize()

	return fyne.NewSize(fyne.Max(titleMinSz.Width, bodyMinSz.Width), titleMinSz.Height+bodyMinSz.Height)
}

func (l *LuckyDrawLayout) Layout(_ []fyne.CanvasObject, containerSize fyne.Size) {
	l.title.Move(fyne.NewPos(0, 0))
	l.title.Resize(l.title.MinSize())

	l.body.Move(fyne.NewPos(0, l.title.Size().Height))
	l.body.Resize(fyne.NewSize(l.body.MinSize().Width, containerSize.Height-l.title.Size().Height))
}

func NewLuckyDrawLayout(title fyne.CanvasObject, body fyne.CanvasObject) fyne.Layout {
	return &LuckyDrawLayout{
		title: title,
		body:  body,
	}
}
