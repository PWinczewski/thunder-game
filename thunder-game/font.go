package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var Fonts map[string]*text.GoTextFaceSource

func loadFont(path string) (*text.GoTextFaceSource, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	face, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return face, nil
}

func LoadFonts() {
	Fonts = make(map[string]*text.GoTextFaceSource)

	err := filepath.Walk("./assets/fonts", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fnt, err := loadFont(path)
		if err != nil {
			log.Fatal(err)
		}
		name := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
		Fonts[name] = fnt

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
