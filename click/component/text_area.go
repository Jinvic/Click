package component

import (
	"fmt"
	"image/color"

	"github.com/Jinvic/Click/click/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

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
			util.NewTextFace(nil, util.DefaultFontSize),
			util.NewCenterDrawOption(width, height))
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
		util.NewTextFace(nil, util.DefaultFontSize),
		util.NewCenterDrawOption(t.image.Bounds().Dx(), t.image.Bounds().Dy()))
	fmt.Println(t.text)
}
