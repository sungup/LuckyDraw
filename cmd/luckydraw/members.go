package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Members struct {
	title   string
	members []string
	widget  *widget.List
}

func (m *Members) Len() int {
	return len(m.members)
}

func (m *Members) CreateItem() fyne.CanvasObject {
	return widget.NewLabel(m.title)
}

func (m *Members) UpdateItem(i widget.ListItemID, o fyne.CanvasObject) {
	o.(*widget.Label).SetText(m.members[i])
}

func (m *Members) Widget(width, height float32) fyne.CanvasObject {
	c := container.NewWithoutLayout()

	label := widget.NewLabel(m.title)
	label.Resize(fyne.NewSize(width, 32.0))
	m.widget.Resize(fyne.NewSize(width, height))
	m.widget.Move(fyne.NewPos(0.0, 32.0))
	c.Add(label)
	c.Add(m.widget)

	c.Resize(fyne.NewSize(width, height))
	return c
}

func NewMembers(title string, members ...string) *Members {
	m := &Members{
		title:   title,
		members: members,
		widget:  nil,
	}

	m.widget = widget.NewList(m.Len, m.CreateItem, m.UpdateItem)

	return m
}
