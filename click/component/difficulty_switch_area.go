package component

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type DifficultySwitchArea struct {
	ComponentBasic
	image *ebiten.Image

	HintTextArea        *TextArea
	DifficultySelectBox *SelectBox
	// valueSelectBox      *SelectBox
}

// TODO: 添加自定义难度
func NewDifficultySwitchArea(x, y, width, height int) *DifficultySwitchArea {
	image := ebiten.NewImage(width, height)
	image.Fill(color.Gray{Y: 128})
	componentBasic := NewComponentBasic(x, y, width, height)
	difficultySwitchArea := DifficultySwitchArea{
		ComponentBasic: *componentBasic,
		image:          image,
	}

	hintTextArea := NewTextArea(0, int(DefaultLineHeight), width, height, "Choose difficulty")
	hintTextArea.SetOption(TextAreaOptionUp)

	selectBoxWidth := 120
	selectBoxHeight := int(DefaultLineHeight * 3)
	selectBoxX := (width - selectBoxWidth) / 2
	selectBoxY := (height - selectBoxHeight) / 2
	difficultySelectBox := NewSelectBox(selectBoxX, selectBoxY, selectBoxWidth, selectBoxHeight, SelectTypeSingle)
	difficultySelectBox.SetOptions([]string{
		string(GameDifficultyNameEasy),
		string(GameDifficultyNameMedium),
		string(GameDifficultyNameHard),
		// string(GameDifficultyNameCustom),
	})
	// 默认选择Medium
	difficultySelectBox.Select(1)
	difficultySelectBox.Choose(1)

	hintTextArea.father = &difficultySwitchArea
	difficultySelectBox.father = &difficultySwitchArea
	difficultySwitchArea.HintTextArea = hintTextArea
	difficultySwitchArea.DifficultySelectBox = difficultySelectBox
	return &difficultySwitchArea
}

func (d *DifficultySwitchArea) Draw(screen *ebiten.Image) {
	d.HintTextArea.Draw(d.image)
	d.DifficultySelectBox.Draw(d.image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(d.x), float64(d.y))
	screen.DrawImage(d.image, op)
}
