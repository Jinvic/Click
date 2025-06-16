package component

import (
	"github.com/Jinvic/Click/click/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 定义字符集类型
type CharSet int

const (
	CharSetNumbers   CharSet = 1 << iota // 数字
	CharSetLowercase                     // 小写字母
	CharSetUppercase                     // 大写字母
	CharSetSymbols                       // 符号集
)

type TextInputBox struct {
	TextArea
	text    string
	counter int
	limit   int

	allowedChars map[rune]bool
	charSet      CharSet // 当前字符集

	onConfirm func()
	onCancel  func()
}

func NewTextInputBox(x, y, width, height int) *TextInputBox {
	box := &TextInputBox{
		TextArea: *NewTextArea(x, y, width, height, ""),
		text:     "",
		counter:  0,
		limit:    50,
	}
	box.SetCharSet(CharSetNumbers | CharSetLowercase | CharSetUppercase)
	return box
}

func (t *TextInputBox) GetText() string {
	return t.text
}

func (t *TextInputBox) SetText(text string) {
	t.text = text
}

func (t *TextInputBox) SetLimit(limit int) {
	t.limit = limit
}

// 设置字符集
func (t *TextInputBox) SetCharSet(charSet CharSet) {
	t.charSet = charSet
	t.updateAllowedChars()
}

// 更新允许的字符集
func (t *TextInputBox) updateAllowedChars() {
	t.allowedChars = make(map[rune]bool)
	if t.charSet == 0 {
		return // 无字符集限制，允许所有字符
	}

	// 根据字符集选项添加允许的字符
	if t.charSet&CharSetNumbers != 0 {
		for _, char := range "0123456789" {
			t.allowedChars[char] = true
		}
	}
	if t.charSet&CharSetLowercase != 0 {
		for _, char := range "abcdefghijklmnopqrstuvwxyz" {
			t.allowedChars[char] = true
		}
	}
	if t.charSet&CharSetUppercase != 0 {
		for _, char := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			t.allowedChars[char] = true
		}
	}
	if t.charSet&CharSetSymbols != 0 {
		for _, char := range "!@#$%^&*()_+-=[]{}|;':\",./<>?~`" {
			t.allowedChars[char] = true
		}
	}
}

func (t *TextInputBox) SetOnConfirm(onConfirm func()) {
	t.onConfirm = onConfirm
}

func (t *TextInputBox) SetOnCancel(onCancel func()) {
	t.onCancel = onCancel
}

func (t *TextInputBox) Update() error {
	t.counter = (t.counter + 1) % 60

	// 输入文本内容
	runes := ebiten.AppendInputChars(nil)
	if len(runes) > 0 {
		text := t.text
		if len(text) < t.limit {
			// 检查字符是否在允许的字符集中
			if len(t.allowedChars) == 0 || t.allowedChars[runes[0]] {
				t.text = text + string(runes[0])
			}
		}
	}

	// 按下退格键，删除字符
	if util.IsKeyLongPressed(ebiten.KeyBackspace) {
		if len(t.text) > 0 {
			t.text = t.text[:len(t.text)-1]
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		t.onConfirm()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		t.onCancel()
	}
	return nil
}

func (t *TextInputBox) Draw(screen *ebiten.Image) {
	text := t.text
	if t.counter%60 < 30 {
		text = text + "_"
	}
	t.TextArea.UpdateText(text)
	t.TextArea.Draw(screen)
}
