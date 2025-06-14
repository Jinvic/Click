package component

import (
	"image/color"

	"github.com/Jinvic/Click/click/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Button struct {
	ComponentBasic
	image *ebiten.Image
}

// 默认按钮大小
const (
	BUTTON_WIDTH  = 100
	BUTTON_HEIGHT = 50
)

func NewButton(x, y, width, height int, str string) *Button {
	image := ebiten.NewImage(width, height)
	image.Fill(color.Gray{Y: 128})
	if len(str) > 0 {
		// ebitenutil.DebugPrint(image, str)
		vector.StrokeRect(image, 0, 0, float32(width), float32(height), 5, color.White, false)
		text.Draw(image,
			str,
			util.NewTextFace(nil, util.DefaultFontSize),
			util.NewCenterDrawOption(width, height))
	}
	return &Button{
		ComponentBasic: *NewComponentBasic(x, y, width, height),
		image:          image,
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.image, op)
}
