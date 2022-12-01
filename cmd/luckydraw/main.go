package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"os"
)

func NewDrawTable(total int, drawCount []int) []*Members {
	drawTable := make([]*Members, 0)
	drawTable = append(drawTable, NewMembers("AllMembers", total))

	for i, cnt := range drawCount {
		drawTable = append(drawTable, NewMembers(fmt.Sprintf("Top %d", i+1), cnt))
	}

	return drawTable
}

func main() {
	_ = os.Setenv("FYNE_FONT", "fonts/BMJUA_ttf.ttf")

	a := app.New()
	w := a.NewWindow("Lucky Draw")

	drawTable := NewDrawTable(32, []int{1, 2, 3, 4, 5, 7, 10})
	info := NewInformation()

	openFile := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
		if err == nil && closer != nil {
			if err = info.Load(closer); err != nil {
				panic(err)
			}

			drawTable[0].Reset(info.Members...)
			for i := 1; i < len(drawTable); i++ {
				drawTable[i].Reset()
			}
		}
	}, w)
	openFile.SetFilter(storage.NewExtensionFileFilter([]string{".yaml", ".yml"}))

	title := container.NewHBox(
		info.Widget(),
		container.NewVBox(
			widget.NewButton("Open File", func() { openFile.Show() }),
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
