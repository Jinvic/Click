package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Component interface {
	Position() (x, y int)
	Draw(screen *ebiten.Image)
}

type TextArea struct {
	x     int
	y     int
	text  string
	image *ebiten.Image
}

func NewTextArea(x, y, width, height int, str string) *TextArea {
	image := ebiten.NewImage(width, height)
	image.Fill(color.Gray{Y: 128})
	if len(str) > 0 {
		// ebitenutil.DebugPrint(image, text)
		text.Draw(image,
			str,
			newTextFace(nil, defaultFontSize),
			newCenterDrawOption(width, height))
	}
	return &TextArea{x: x, y: y, text: str, image: image}
}

func (t *TextArea) Position() (x, y int) {
	return t.x, t.y
}

func (t *TextArea) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.x), float64(t.y))
	screen.DrawImage(t.image, op)
}

func (t *TextArea) UpdateText(str string) {
	t.image.Fill(color.Gray{Y: 128})
	t.text = str
	// ebitenutil.DebugPrint(t.image, t.text)
	text.Draw(t.image,
		t.text,
		newTextFace(nil, defaultFontSize),
		newCenterDrawOption(t.image.Bounds().Dx(), t.image.Bounds().Dy()))
	fmt.Println(t.text)
}

type Button struct {
	x      int
	y      int
	width  int
	height int
	image  *ebiten.Image
}

func NewButton(x, y, width, height int, str string) *Button {
	image := ebiten.NewImage(width, height)
	image.Fill(color.Gray{Y: 128})
	if len(str) > 0 {
		// ebitenutil.DebugPrint(image, str)
		text.Draw(image,
			str,
			newTextFace(nil, defaultFontSize),
			newCenterDrawOption(width, height))
	}
	return &Button{x: x, y: y, width: width, height: height, image: image}
}

func (b *Button) Position() (x, y int) {
	return b.x, b.y
}

func (b *Button) IsButtonJustPressed() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if mx >= b.x && mx <= b.x+b.width && my >= b.y && my <= b.y+b.height {
			return true
		}
	}
	return false
}

func (b *Button) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.image, op)
}
