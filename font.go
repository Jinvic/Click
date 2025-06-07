package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	defaultFaceSource *text.GoTextFaceSource
)

const (
	defaultFontSize = 16
)

func init() {
	// init FaceSource
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	defaultFaceSource = s
}

func newCenterDrawOption(width, height int) *text.DrawOptions {
	drawOptions := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign:   text.AlignCenter, // 水平居中
			SecondaryAlign: text.AlignCenter, // 垂直居中
		},
	}
	drawOptions.GeoM.Translate(float64(width/2), float64(height/2))
	return drawOptions
}

func newTextFace(faceSource *text.GoTextFaceSource, size int) *text.GoTextFace {
	if faceSource == nil {
		faceSource = defaultFaceSource
	}
	return &text.GoTextFace{
		Source: faceSource,
		Size:   float64(size),
	}
}
