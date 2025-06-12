package component

import (
	"image/color"

	"github.com/Jinvic/Click/click/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type SelectType string

const (
	SelectTypeSingle SelectType = "Single"
	SelectTypeMulti  SelectType = "Multi"
)

const (
	SelectCursorTextSingleSelected   = "-->" // 单选选中
	SelectCursorTextSingleUnselected = "   " // 单选未选中
	SelectCursorTextMultiSelected    = "[x]" // 多选选中
	SelectCursorTextMultiUnselected  = "[ ]" // 多选未选中
)

// TODO
type SelectCursor struct {
	ComponentBasic
	text string
}

func NewSelectCursor(x, y, width, height int, selectType SelectType) *SelectCursor {
	var text string
	switch selectType {
	case SelectTypeSingle:
		text = SelectCursorTextSingleUnselected
	case SelectTypeMulti:
		text = SelectCursorTextMultiUnselected
	}

	selectCursor := SelectCursor{
		ComponentBasic: *NewComponentBasic(x, y, width, height),
		text:           text,
	}
	return &selectCursor
}

type SelectOption struct {
	ComponentBasic
	TextArea
	text   string
	cursor *SelectCursor
}

func NewSelectOption(x, y, width, height int, optionText string, selectType SelectType) *SelectOption {
	option := SelectOption{
		text:           optionText,
		cursor:         NewSelectCursor(x, y, width, height, selectType),
		ComponentBasic: *NewComponentBasic(x, y, width, height),
		TextArea:       *NewTextArea(x, y, width, height, optionText),
	}
	return &option
}

func (s *SelectOption) Draw(screen *ebiten.Image) {
	text := s.cursor.text + s.text
	s.TextArea.UpdateText(text)
	s.TextArea.Draw(screen)
}

func (s *SelectOption) Select(selectType SelectType) {
	if selectType == SelectTypeSingle {
		s.cursor.text = SelectCursorTextSingleSelected
	} else {
		s.cursor.text = SelectCursorTextMultiSelected
	}
}

func (s *SelectOption) Unselect(selectType SelectType) {
	if selectType == SelectTypeSingle {
		s.cursor.text = SelectCursorTextSingleUnselected
	} else {
		s.cursor.text = SelectCursorTextMultiUnselected
	}
}

type SelectBox struct {
	ComponentBasic
	image *ebiten.Image

	selectType SelectType
	options    []SelectOption
	selected   util.Set[int]
}

func NewSelectBox(x, y, width, height int, selectType SelectType) *SelectBox {
	image := ebiten.NewImage(width, height)
	image.Fill(color.Gray{Y: 128})
	selectBox := SelectBox{
		ComponentBasic: *NewComponentBasic(x, y, width, height),
		image:          image,
		selectType:     selectType,
		selected:       util.Set[int]{},
	}
	return &selectBox
}

func (s *SelectBox) SetOptions(optionTexts []string) {
	// 计算行高
	lineHeight := float64(s.height) / float64(len(optionTexts))
	var options []SelectOption
	for i, optionText := range optionTexts {
		option := NewSelectOption(0, int(float64(i)*lineHeight), s.width, int(lineHeight), optionText, s.selectType)
		options = append(options, *option)
		option.father = s
	}
	s.options = options

}

func (s *SelectBox) Draw(screen *ebiten.Image) {
	for _, option := range s.options {
		option.Draw(s.image)
	}
}

func (s *SelectBox) Select(index int) {
	if s.selectType == SelectTypeSingle {
		for selectedIndex := range s.selected {
			s.options[selectedIndex].Unselect(s.selectType)
		}
		s.selected.Clear()
		s.selected.Add(index)
		s.options[index].Select(s.selectType)
	} else {
		if s.selected.Contains(index) {
			s.selected.Remove(index)
			s.options[index].Unselect(s.selectType)
		} else {
			s.selected.Add(index)
			s.options[index].Select(s.selectType)
		}
	}
}
