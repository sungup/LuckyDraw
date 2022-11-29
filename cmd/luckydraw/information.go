package main

import (
	"gopkg.in/yaml.v3"
	"io"
)

type Information struct {
	Title   string   `yaml:"title"`
	Intro   string   `yaml:"intro"`
	Message string   `yaml:"message"`
	Members []string `yaml:"members"`
}

func LoadInformation(reader io.Reader) (info *Information, err error) {
	var (
		buf []byte
	)

	if buf, err = io.ReadAll(reader); err != nil {
		return info, err
	}

	info = &Information{}

	return info, yaml.Unmarshal(buf, info)
}
