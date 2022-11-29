package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
)

const (
	top1Members = 1
	top2Members = 2
	top3Members = 3
	top4Members = 4
	top5Members = 5
	top6Members = 7
	top7Members = 10
)

var (
	memberCount = []int{32, 1, 2, 3, 4, 5, 7, 10}
)

type LuckyDrawLayout struct {
	title fyne.CanvasObject
	body  fyne.CanvasObject
}

func (l *LuckyDrawLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	titleMinSz := l.title.MinSize()
	bodyMinSz := l.body.MinSize()

	return fyne.NewSize(fyne.Max(titleMinSz.Width, bodyMinSz.Width), titleMinSz.Height+bodyMinSz.Height)
}

func (l *LuckyDrawLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
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

func main() {
	_ = os.Setenv("FYNE_FONT", "fonts/BMJUA_ttf.ttf")

	info, err := LoadInformation("luckydraw.yaml")
	if err != nil {
		panic(err)
	}

	a := app.New()
	w := a.NewWindow("Lucky Draw")

	title := container.NewHBox(
		widget.NewCard(info.Title, info.Intro, widget.NewRichTextWithText(info.Message)),
		container.NewVBox(
			widget.NewButton("Shuffle", func() {}),
			widget.NewButton("Go!", func() {}),
		),
	)

	partMembers := NewMembers("All Members", info.Members...)
	top1 := NewMembers("Top 1")
	top2 := NewMembers("Top 2")
	top3 := NewMembers("Top 3")
	top4 := NewMembers("Top 4")
	top5 := NewMembers("Top 5")
	top6 := NewMembers("Top 6")
	top7 := NewMembers("Top 7")

	body := container.NewHBox(
		partMembers.Widget(200, 600),
		container.NewGridWithRows(
			3,
			top1.Widget(200, 600),
			top2.Widget(200, 600),
			top3.Widget(200, 600),
		),
		container.NewGridWithRows(
			2,
			top4.Widget(200, 600),
			top5.Widget(200, 600),
		),
		top6.Widget(200, 600),
		top7.Widget(200, 600),
	)

	body.Resize(fyne.NewSize(0, 0))

	doc := container.New(NewLuckyDrawLayout(title, body), title, body)

	w.SetContent(doc)
	w.Resize(fyne.Size{Width: 1024, Height: 768})

	w.ShowAndRun()

}
