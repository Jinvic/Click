package util

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	DefaultFaceSource *text.GoTextFaceSource
)

const (
	DefaultFontSize = 24
)

func init() {
	// init FaceSource
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	DefaultFaceSource = s
}

// 水平居中,垂直居中
func NewCenterDrawOption(width, height int) *text.DrawOptions {
	drawOptions := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign:   text.AlignCenter, // 水平居中
			SecondaryAlign: text.AlignCenter, // 垂直居中
		},
	}
	drawOptions.GeoM.Translate(float64(width/2), float64(height/2))
	return drawOptions
}

// 水平居左,垂直居中
func NewLeftDrawOption(width, height int) *text.DrawOptions {
	drawOptions := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign:   text.AlignStart, // 水平居左
			SecondaryAlign: text.AlignCenter, // 垂直居中
		},
	}
	drawOptions.GeoM.Translate(0, float64(height/2))
	return drawOptions
}

// 水平居中,垂直居上
func NewUpDrawOption(width, height int) *text.DrawOptions {
	drawOptions := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign:   text.AlignCenter, // 水平居中
			SecondaryAlign: text.AlignStart, // 垂直居上
		},
	}
	drawOptions.GeoM.Translate(float64(width/2), 0)
	return drawOptions
}


// 水平居中,高度指定
func NewHCenterDrawOption(width, height int, y float64) *text.DrawOptions {
	drawOptions := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign:   text.AlignCenter, // 水平居中
			SecondaryAlign: text.AlignStart,
		},
	}
	drawOptions.GeoM.Translate(float64(width/2), y)
	return drawOptions
}

// 水平居左,高度指定
func NewHLeftDrawOption(width, height int, y float64) *text.DrawOptions {
	drawOptions := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			PrimaryAlign:   text.AlignStart,
			SecondaryAlign: text.AlignStart,
		},
	}
	drawOptions.GeoM.Translate(0, y)
	return drawOptions
}

func NewTextFace(faceSource *text.GoTextFaceSource, size int) *text.GoTextFace {
	if faceSource == nil {
		faceSource = DefaultFaceSource
	}
	return &text.GoTextFace{
		Source: faceSource,
		Size:   float64(size),
	}
}
