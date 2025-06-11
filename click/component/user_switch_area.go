package component

import "github.com/hajimehoshi/ebiten/v2"

type UserSwitchArea struct {
	MultiTextArea
	username string
	cursorCounter  int
}

func NewUserSwitchArea(x, y, width, height int, username string) *UserSwitchArea {
	area := UserSwitchArea{username: username}
	area.MultiTextArea = *NewMultiTextArea(x, y, width, height, []string{
		"Enter username(max 10 characters).",
		"Press Enter to confirm.",
		"Press ESC to cancel.",
		username})
	return &area
}

func (u *UserSwitchArea) GetUsername() string {
	return u.username
}

func (u *UserSwitchArea) SetUsername(username string) {
	u.username = username
}

// 光标闪烁计数器
func (u *UserSwitchArea) UpdateCursorCounter() {
	u.cursorCounter = (u.cursorCounter + 1) % 60
}

func (u *UserSwitchArea) Draw(screen *ebiten.Image) {
	username := u.username
	if u.cursorCounter%60 < 30 { // 闪烁光标
		username += "_"
	}
	texts := u.GetTexts()
	texts[len(texts)-1] = username
	u.UpdateTexts(texts)
	u.MultiTextArea.Draw(screen)
}
