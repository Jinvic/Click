package component

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type UserSwitchArea struct {
	ComponentBasic
	image *ebiten.Image

	multiTextArea *MultiTextArea
	textInputBox *TextInputBox
}

func NewUserSwitchArea(x, y, width, height int, username string) *UserSwitchArea {
	image := ebiten.NewImage(width, height)
	image.Fill(color.Gray{Y: 128})
	componentBasic := NewComponentBasic(x, y, width, height)
	area := UserSwitchArea{
		ComponentBasic: *componentBasic,
		image:          image,
	}
	multiTextArea := NewMultiTextArea(0, 0, width, height-2*int(DefaultLineHeight), []string{
		"Enter username(max 10 characters).",
		"Press Enter to confirm.",
		"Press ESC to cancel."})
	textInputBox := NewTextInputBox(x, height-2*int(DefaultLineHeight), width, 2*int(DefaultLineHeight), username, 10)
	multiTextArea.father = &area
	textInputBox.father = &area
	area.multiTextArea = multiTextArea
	area.textInputBox = textInputBox
	return &area
}

func (u *UserSwitchArea) GetUsername() string {
	return u.textInputBox.text
}

func (u *UserSwitchArea) SetUsername(username string) {
	u.textInputBox.text = username
}

// 光标闪烁计数器
func (u *UserSwitchArea) Update() error {
	return u.textInputBox.Update()
}

func (u *UserSwitchArea) Draw(screen *ebiten.Image) {
	u.image.Fill(color.Gray{Y: 128})
	u.multiTextArea.Draw(u.image)
	u.textInputBox.Draw(u.image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(u.x), float64(u.y))
	screen.DrawImage(u.image, op)
}
