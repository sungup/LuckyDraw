package main

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Information struct {
	Title   string   `yaml:"title"`
	Intro   string   `yaml:"intro"`
	Message string   `yaml:"message"`
	Members []string `yaml:"members"`
}

func LoadInformation(filepath string) (info *Information, err error) {
	var (
		fp  *os.File
		buf []byte
	)
	if fp, err = os.Open(filepath); err != nil {
		return info, err
	}

	if buf, err = io.ReadAll(fp); err != nil {
		return info, err
	}

	info = &Information{}

	return info, yaml.Unmarshal(buf, info)
}
