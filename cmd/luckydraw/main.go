package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	partMembers := NewMembers("All Members", "Sungup", "Lucy", "Seonho", "Dan", "Changduk")
	top1 := NewMembers("Top 1", "hello")
	top2 := NewMembers("Top 2")
	top3 := NewMembers("Top 3")
	top4 := NewMembers("Top 4")
	top5 := NewMembers("Top 5")
	top6 := NewMembers("Top 6")
	top7 := NewMembers("Top 7")

	hello := widget.NewLabel("Helly Fyne!")
	button := widget.NewButton(
		"Hi!",
		func() {
			hello.SetText("Welcome :)")
		},
	)

	body := container.NewGridWithColumns(4,
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

	left := partMembers.Widget(200, 600)

	dummy := container.NewWithoutLayout()
	dummy.Resize(fyne.NewSize(1, 1))
	doc := container.NewBorder(hello, button, left, body)
	doc.Add(hello)
	doc.Add(button)
	doc.Add(body)
	doc.Add(left)

	container.NewVBox()

	w.SetContent(doc)
	w.Resize(fyne.Size{Width: 1024, Height: 768})

	w.ShowAndRun()

}
