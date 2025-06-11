package component

import (
	"image/color"

	"github.com/Jinvic/Click/click/log"
	"github.com/Jinvic/Click/click/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// 单行文本
type TextArea struct {
	ComponentBasic
	text   string
	image  *ebiten.Image
}

func NewTextArea(x, y, width, height int, str string) *TextArea {
	image := ebiten.NewImage(width, height)
	area := TextArea{
		ComponentBasic: *NewComponentBasic(x, y, width, height),
		text:           str,
		image:          image,
	}
	area.UpdateText(str)
	return &area
}

func (t *TextArea) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.x), float64(t.y))
	screen.DrawImage(t.image, op)
}

func (t *TextArea) UpdateText(str string) {
	t.image.Fill(color.Gray{Y: 128})
	t.text = str
	text.Draw(t.image,
		t.text,
		util.NewTextFace(nil, util.DefaultFontSize),
		util.NewCenterDrawOption(t.width, t.height))
	log.Debug(t.text)
}

func (t *TextArea) GetText() string {
	return t.text
}

// 多行文本
type MultiTextArea struct {
	ComponentBasic
	texts  []string
	image  *ebiten.Image
}

func NewMultiTextArea(x, y, width, height int, strs []string) *MultiTextArea {
	image := ebiten.NewImage(width, height)
	area := MultiTextArea{
		ComponentBasic: *NewComponentBasic(x, y, width, height),
		texts:          strs,
		image:          image,
	}
	area.UpdateTexts(strs)
	return &area
}

func (t *MultiTextArea) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.x), float64(t.y))
	screen.DrawImage(t.image, op)
}

func (t *MultiTextArea) UpdateTexts(strs []string) {
	t.image.Fill(color.Gray{Y: 128})
	t.texts = strs

	face := util.NewTextFace(nil, util.DefaultFontSize)
	// 计算行高
	// lineHeight := face.Metrics().HAscent
	lineHeight := float64(t.height) / float64(len(strs)+2)

	// 逐行绘制文本
	for i, str := range strs {
		y := lineHeight + (lineHeight * float64(i)) // 每行向下偏移一个行高
		text.Draw(t.image, str, face, util.NewHLeftDrawOption(t.width, t.height, y))
	}
}

func (t *MultiTextArea) GetTexts() []string {
	return t.texts
}
