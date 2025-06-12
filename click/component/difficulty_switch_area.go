package component

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type DifficultySwitchArea struct {
	ComponentBasic
	image *ebiten.Image

	// hintTextArea     *TextArea
	// difficultySelect *SelectBox
	// valueSelect      *SelectBox
}

func NewDifficultySwitchArea(x, y, width, height int) *DifficultySwitchArea {
	image := ebiten.NewImage(width, height)
	image.Fill(color.Gray{Y: 128})
	componentBasic := NewComponentBasic(x, y, width, height)
	difficultySwitchArea := DifficultySwitchArea{
		ComponentBasic: *componentBasic,
		image:          image,
	}
	return &difficultySwitchArea
}
