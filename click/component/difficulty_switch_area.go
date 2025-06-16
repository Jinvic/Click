package component

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/cast"
)

const (
	DifficultySwitchAreaStatusDifficulty = iota // 0: 选择难度
	DifficultySwitchAreaStatusValue             // 1: 选择值
	DifficultySwitchAreaStatusValueInput        // 2: 修改值
)

type DifficultySwitchArea struct {
	ComponentBasic
	image *ebiten.Image

	status     int
	difficulty GameDifficulty

	HintTextArea        *TextArea
	DifficultySelectBox *SelectBox
	ValueSelectBox      *SelectBox
	ValueInputBox       *TextInputBox
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

	selectBoxWidth := 160
	selectBoxHeight := int(DefaultLineHeight * 4)
	selectBoxX := (width - selectBoxWidth) / 2
	selectBoxY := (height - selectBoxHeight) / 2

	difficultySelectBox := NewSelectBox(selectBoxX, selectBoxY, selectBoxWidth, selectBoxHeight, SelectTypeSingle)
	difficultySelectBox.SetOptions([]string{
		string(GameDifficultyNameEasy),
		string(GameDifficultyNameMedium),
		string(GameDifficultyNameHard),
		string(GameDifficultyNameCustom),
	})
	// 默认选择Medium
	difficultySelectBox.Select(1)
	difficultySelectBox.Choose(1)

	valueSelectBox := NewSelectBox(selectBoxX, selectBoxY, selectBoxWidth*2, selectBoxHeight, SelectTypeSingle)

	valueInputBox := NewTextInputBox((width-selectBoxWidth)/2, height-int(DefaultLineHeight*2), selectBoxWidth, int(DefaultLineHeight))
	valueInputBox.SetText("")
	valueInputBox.SetLimit(5)
	valueInputBox.SetCharSet(CharSetNumbers)
	valueInputBox.SetOption(TextAreaOptionLeft)

	hintTextArea.father = &difficultySwitchArea
	difficultySelectBox.father = &difficultySwitchArea
	valueSelectBox.father = &difficultySwitchArea
	valueInputBox.father = &difficultySwitchArea
	difficultySwitchArea.HintTextArea = hintTextArea
	difficultySwitchArea.DifficultySelectBox = difficultySelectBox
	difficultySwitchArea.ValueSelectBox = valueSelectBox
	difficultySwitchArea.ValueInputBox = valueInputBox

	difficultySwitchArea.status = DifficultySwitchAreaStatusDifficulty
	difficultySwitchArea.SetDifficulty(GameDifficultyMedium) // 默认选择Medium
	return &difficultySwitchArea
}

func (d *DifficultySwitchArea) Draw(screen *ebiten.Image) {
	d.HintTextArea.Draw(d.image)
	switch d.status {
	case DifficultySwitchAreaStatusDifficulty:
		d.DifficultySelectBox.Draw(d.image)
	case DifficultySwitchAreaStatusValue:
		d.ValueSelectBox.Draw(d.image)
	case DifficultySwitchAreaStatusValueInput:
		d.ValueSelectBox.Draw(d.image)
		d.ValueInputBox.Draw(d.image)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(d.x), float64(d.y))
	screen.DrawImage(d.image, op)
}

func (d *DifficultySwitchArea) GetStatus() int {
	return d.status
}

func (d *DifficultySwitchArea) SetStatus(status int) {
	d.status = status
}

func (d *DifficultySwitchArea) SetDifficulty(difficulty GameDifficulty) {
	d.difficulty = difficulty

	difficultyText := fmt.Sprintf(difficultyTextTemplate, "", difficulty.Radius, difficulty.Speed, difficulty.Duration)
	valueOptions := strings.Split(difficultyText, "\n")[1:]
	d.ValueSelectBox.SetOptions(valueOptions)
	d.ValueSelectBox.Select(0)
	d.ValueSelectBox.Choose(0)
}

func (d *DifficultySwitchArea) GetValue() int {
	var selectedIndex int
	selected := d.ValueSelectBox.GetSelected()
	if len(selected) == 0 {
		selectedIndex = 0
	} else {
		selectedIndex = selected[0]
	}

	return d.GetValueByIndex(selectedIndex)
}

func (d *DifficultySwitchArea) GetValueByIndex(index int) int {
	optionText := d.ValueSelectBox.options[index].text
	valueStr := strings.Split(optionText, ":")[1]
	value := cast.ToInt(valueStr)
	return value
}

func (d *DifficultySwitchArea) SetValue() {
	var selectedIndex int
	selected := d.ValueSelectBox.GetSelected()
	if len(selected) == 0 {
		selectedIndex = 0
	} else {
		selectedIndex = selected[0]
	}

	valueStr := d.ValueInputBox.GetText()
	trimmed := strings.TrimLeft(valueStr, "0")
	if trimmed == "" {
		trimmed = "0"
	}
	value := cast.ToInt(trimmed)
	if value == 0 {
		return
	}

	optionText := d.ValueSelectBox.options[selectedIndex].text
	optionText = strings.Split(optionText, ":")[0] + ":" + cast.ToString(value)
	d.ValueSelectBox.options[selectedIndex].text = optionText
}

func (d *DifficultySwitchArea) GetCustomDifficulty() GameDifficulty {
	radius := d.GetValueByIndex(0)
	speed := d.GetValueByIndex(1)
	duration := d.GetValueByIndex(2)
	difficulty := GameDifficulty{
		Name:     GameDifficultyNameCustom,
		Radius:   radius,
		Speed:    speed,
		Duration: duration,
	}
	return difficulty
}

func (d *DifficultySwitchArea) SwitchStatus(status int) {
	switch status {
	case DifficultySwitchAreaStatusDifficulty:
		d.HintTextArea.UpdateText("Choose difficulty")
		d.SetStatus(DifficultySwitchAreaStatusDifficulty)
	case DifficultySwitchAreaStatusValue:
		d.HintTextArea.UpdateText("Choose value")
		d.SetStatus(DifficultySwitchAreaStatusValue)
	case DifficultySwitchAreaStatusValueInput:
		d.HintTextArea.UpdateText("Input value")
		d.SetStatus(DifficultySwitchAreaStatusValueInput)
	}
}