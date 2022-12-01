package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v3"
	"io"
)

type Information struct {
	Title   string   `yaml:"title"`
	Intro   string   `yaml:"intro"`
	Message string   `yaml:"message"`
	Members []string `yaml:"members"`

	card *widget.Card
}

const (
	defaultTitle    = "Hello Lucky Draw!!"
	defaultSubTitle = "Lucky Draw Program for the Special Event"
	defaultMessage  = "Before start shuffle and draw, please open the event YAML file first!"
)

func NewInformation() *Information {
	return &Information{
		Title:   defaultTitle,
		Intro:   defaultSubTitle,
		Message: defaultMessage,
		Members: nil,
		card:    widget.NewCard(defaultTitle, defaultSubTitle, widget.NewRichTextWithText(defaultMessage)),
	}
}

func (i *Information) Widget() fyne.CanvasObject {
	return i.card
}

func (i *Information) Load(reader io.Reader) (err error) {
	var buf []byte

	if buf, err = io.ReadAll(reader); err == nil {
		if err = yaml.Unmarshal(buf, i); err == nil {
			i.card.SetTitle(i.Title)
			i.card.SetSubTitle(i.Intro)
			i.card.Content.(*widget.RichText).ParseMarkdown(i.Message)
		}
	}

	i.card.Refresh()

	return err
}
