package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"time"
)

var (
	rnd *rand.Rand
)

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Members struct {
	title   string
	members []string
	max     int
	widget  *widget.List
	card    *widget.Card
}

func (m *Members) subtitle() string {
	return fmt.Sprintf("Current Member: %d / Max: %d", len(m.members), m.max)
}

func (m *Members) Len() int                      { return len(m.members) }
func (m *Members) Max() int                      { return m.max }
func (m *Members) IsFull() bool                  { return m.Len() == m.Max() }
func (m *Members) Swap(i, j int)                 { m.members[i], m.members[j] = m.members[j], m.members[i] }
func (m *Members) CreateItem() fyne.CanvasObject { return widget.NewLabel(m.title) }

func (m *Members) UpdateItem(i widget.ListItemID, o fyne.CanvasObject) {
	o.(*widget.Label).SetText(m.members[i])
}

func (m *Members) Refresh() {
	m.card.SetSubTitle(m.subtitle())
	m.card.Refresh()
	m.widget.Refresh()
}

func (m *Members) Shuffle() {
	rnd.Shuffle(m.Len(), m.Swap)

	m.Refresh()
}

func (m *Members) Reset(name ...string) {
	m.members = name

	m.Refresh()
}

func (m *Members) Append(name ...string) {
	m.members = append(m.members, name...)

	m.Refresh()
}

func (m *Members) Pick() (name string) {
	pick := rnd.Intn(m.Len())

	name = m.members[pick]
	m.members = append(m.members[0:pick], m.members[pick+1:len(m.members)]...)

	m.Refresh()

	return name
}

func (m *Members) Widget(width, height float32) fyne.CanvasObject {
	m.card.Resize(fyne.NewSize(width, height))
	return m.card
}

func NewMembers(title string, max int, members ...string) *Members {
	m := &Members{
		title:   title,
		members: members,
		max:     max,
		widget:  nil,
	}

	m.widget = widget.NewList(m.Len, m.CreateItem, m.UpdateItem)
	m.card = widget.NewCard(m.title, m.subtitle(), m.widget)

	return m
}
