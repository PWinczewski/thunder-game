package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Sprites map[string]*ebiten.Image

func LoadSprites() {
	Sprites = make(map[string]*ebiten.Image)

	err := filepath.Walk("./assets/art", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Fatal(err)
		}

		name := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
		Sprites[name] = img
		return nil
	})

	fmt.Println(Sprites["tileBurning"])

	if err != nil {
		log.Fatal(err)
	}
}
