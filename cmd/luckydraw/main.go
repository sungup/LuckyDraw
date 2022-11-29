package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
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

func main() {
	_ = os.Setenv("FYNE_FONT", "fonts/BMJUA_ttf.ttf")

	fp, _ := os.Open("luckydraw.yaml")

	info, err := LoadInformation(fp)
	if err != nil {
		panic(err)
	}

	drawTable := make([]*Members, 0)
	for i, cnt := range memberCount {
		var (
			title   string
			members []string
		)

		if i == 0 {
			title = "All Members"
			members = info.Members
		} else {
			title = fmt.Sprintf("Top %d", i)
			members = []string{}
		}

		drawTable = append(drawTable, NewMembers(title, cnt, members...))
	}

	a := app.New()
	w := a.NewWindow("Lucky Draw")

	title := container.NewHBox(
		widget.NewCard(info.Title, info.Intro, widget.NewRichTextWithText(info.Message)),
		container.NewVBox(
			widget.NewButton("Shuffle", drawTable[0].Shuffle),
			widget.NewButton("Go!", func() {
				var name string

				if drawTable[0].Len() > 0 {
					name = drawTable[0].Pick()
				}

				for i := len(drawTable) - 1; i > 0; i-- {
					if !drawTable[i].IsFull() {
						drawTable[i].Append(name)
						break
					}
				}
			}),
		),
	)

	body := container.NewHBox(
		drawTable[0].Widget(200, 600),
		container.NewGridWithRows(
			3,
			drawTable[1].Widget(200, 600),
			drawTable[2].Widget(200, 600),
			drawTable[3].Widget(200, 600),
		),
		container.NewGridWithRows(
			2,
			drawTable[4].Widget(200, 600),
			drawTable[5].Widget(200, 600),
		),
		drawTable[6].Widget(200, 600),
		drawTable[7].Widget(200, 600),
	)

	body.Resize(fyne.NewSize(0, 0))

	doc := container.New(NewLuckyDrawLayout(title, body), title, body)

	w.SetContent(doc)
	w.Resize(fyne.Size{Width: 1024, Height: 768})

	w.ShowAndRun()

}
