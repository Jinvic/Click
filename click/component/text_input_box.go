package component

import (
	"github.com/Jinvic/Click/click/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TextInputBox struct {
	TextArea
	text    string
	counter int
	limit   int

	onConfirm func()
	onCancel  func()
}

func NewTextInputBox(x, y, width, height int, text string, limit int) *TextInputBox {
	box := &TextInputBox{
		TextArea: *NewTextArea(x, y, width, height, text),
		text:     text,
		counter:  0,
		limit:    limit,
	}
	return box
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
			t.text = text + string(runes[0])
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
