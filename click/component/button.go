package component

import (
	"image/color"

	"github.com/Jinvic/Click/click/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

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
			util.NewTextFace(nil, util.DefaultFontSize),
			util.NewCenterDrawOption(width, height))
	}
	return &Button{x: x, y: y, width: width, height: height, image: image}
}

func (b *Button) Position() (x, y int) {
	return b.x, b.y
}

func (b *Button) Size() (width, height int) {
	return b.width, b.height
}

func (b *Button) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.image, op)
}
