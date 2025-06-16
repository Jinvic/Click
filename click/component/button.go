package component

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	TextArea
}

// 默认按钮大小
const (
	BUTTON_WIDTH  = 100
	BUTTON_HEIGHT = 50
)

func NewButton(x, y, width, height int, str string) *Button {
	textArea := NewTextArea(x, y, width, height, str)
	return &Button{
		TextArea: *textArea,
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	vector.StrokeRect(b.image, 0, 0, float32(b.width), float32(b.height), 5, color.White, false)
	b.TextArea.Draw(screen)
}
