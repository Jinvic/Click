package component

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type ConfirmArea struct {
	ComponentBasic
	image *ebiten.Image

	hintTextArea  *TextArea
	confirmButton *Button
	cancelButton  *Button

	onConfirm func()
	onCancel  func()
}

func NewConfirmArea(x, y, width, height int, hintText string) *ConfirmArea {
	image := ebiten.NewImage(width, height)
	image.Fill(color.Gray{Y: 128})
	componentBasic := NewComponentBasic(x, y, width, height)
	confirmArea := ConfirmArea{
		ComponentBasic: *componentBasic,
		image:          image,
	}

	buttonWidth := 60
	hintTextArea := NewTextArea(0, 0, width-buttonWidth, height-BUTTON_HEIGHT, hintText)
	hintTextArea.father = &confirmArea

	confirmButton := NewButton(
		(width-buttonWidth)/4-buttonWidth/2,
		height-BUTTON_HEIGHT,
		buttonWidth,
		BUTTON_HEIGHT,
		"Confirm",
	)
	confirmButton.father = &confirmArea

	cancelButton := NewButton(
		(width-buttonWidth)/4*3-buttonWidth/2,
		height-BUTTON_HEIGHT,
		buttonWidth, BUTTON_HEIGHT,
		"Cancel")
	cancelButton.father = &confirmArea

	confirmArea.hintTextArea = hintTextArea
	confirmArea.confirmButton = confirmButton
	confirmArea.cancelButton = cancelButton
	return &confirmArea
}

func (c *ConfirmArea) SetHintText(hintText string) {
	c.hintTextArea.UpdateText(hintText)
}

func (c *ConfirmArea) SetOnConfirm(onConfirm func()) {
	c.onConfirm = onConfirm
}

func (c *ConfirmArea) SetOnCancel(onCancel func()) {
	c.onCancel = onCancel
}

func (c *ConfirmArea) Draw(screen *ebiten.Image) {
	c.image.Fill(color.Gray{Y: 128})
	c.hintTextArea.Draw(c.image)
	c.confirmButton.Draw(c.image)
	c.cancelButton.Draw(c.image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.x), float64(c.y))
	screen.DrawImage(c.image, op)
}

func (c *ConfirmArea) IsConfirmButtonJustClicked() bool {
	return IsComponentJustClicked(c.confirmButton)
}

func (c *ConfirmArea) IsCancelButtonJustClicked() bool {
	return IsComponentJustClicked(c.cancelButton)
}

func (c *ConfirmArea) OnConfirm() {
	if c.onConfirm != nil {
		c.onConfirm()
	}
}

func (c *ConfirmArea) OnCancel() {
	if c.onCancel != nil {
		c.onCancel()
	}
}
