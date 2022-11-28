package main

import (
	"fyne.io/fyne/v2"
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

func (m *Members) Widget() *widget.List {
	return m.widget
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
