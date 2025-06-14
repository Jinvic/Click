package component

import (
	"image/color"

	"github.com/Jinvic/Click/click/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SelectType string

const (
	SelectTypeSingle SelectType = "Single"
	SelectTypeMulti  SelectType = "Multi"
)

const (
	SelectCursorTextSingleSelected   = ">>>"  // 单选选中
	SelectCursorTextSingleUnselected = "---"  // 单选未选中
	SelectCursorTextMultiSelected    = "-[x]" // 多选选中
	SelectCursorTextMultiUnselected  = "-[ ]" // 多选未选中
)

var (
	SelectCursorTextChoosedFrames map[string][]string
)

func init() {
	SelectCursorTextChoosedFrames = make(map[string][]string)
	SelectCursorTextChoosedFrames["Single"] = []string{
		">>>",
		"->>",
		"-->",
		"---",
		">--",
		">>-",
	}
	SelectCursorTextChoosedFrames["MultiSelected"] = []string{
		">[x]",
		"-[x]",
		">[x]",
		"-[x]",
		">[x]",
		"-[x]",
	}
	SelectCursorTextChoosedFrames["MultiUnselected"] = []string{
		">[ ]",
		"-[ ]",
		">[ ]",
		"-[ ]",
		">[ ]",
		"-[ ]",
	}
}

type SelectCursor struct {
	ComponentBasic
	text       string
	selectType SelectType
	choosed    bool
	selected   bool
	counter    int
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
		selectType:     selectType,
		choosed:        false,
	}
	return &selectCursor
}

func (s *SelectCursor) GetText() string {
	if s.choosed {
		return s.GetChoosedFrame()
	}
	return s.text
}

func (s *SelectCursor) GetChoosedFrame() string {
	// 每30帧切换一次
	frameIndex := s.counter / 30
	if s.selectType == SelectTypeSingle {
		return SelectCursorTextChoosedFrames["Single"][frameIndex%len(SelectCursorTextChoosedFrames["Single"])]
	} else {
		if s.selected {
			return SelectCursorTextChoosedFrames["MultiSelected"][frameIndex%len(SelectCursorTextChoosedFrames["MultiSelected"])]
		} else {
			return SelectCursorTextChoosedFrames["MultiUnselected"][frameIndex%len(SelectCursorTextChoosedFrames["MultiUnselected"])]
		}
	}
}

func (s *SelectCursor) Select() {
	s.selected = true
	if s.selectType == SelectTypeSingle {
		s.text = SelectCursorTextSingleSelected
	} else {
		s.text = SelectCursorTextMultiSelected
	}
}

func (s *SelectCursor) Unselect() {
	s.selected = false
	if s.selectType == SelectTypeSingle {
		s.text = SelectCursorTextSingleUnselected
	} else {
		s.text = SelectCursorTextMultiUnselected
	}
}

func (s *SelectCursor) Choose() {
	s.choosed = true
}

func (s *SelectCursor) Unchoose() {
	s.choosed = false
}

func (s *SelectCursor) UpdateCounter(counter int) {
	s.counter = counter
}

type SelectOption struct {
	ComponentBasic
	TextArea
	text   string
	cursor *SelectCursor
}

func NewSelectOption(x, y, width, height int, optionText string, selectType SelectType) *SelectOption {
	textArea := NewTextArea(x, y, width, height, optionText)
	textArea.SetOption(TextAreaOptionLeft)
	option := SelectOption{
		text:           optionText,
		cursor:         NewSelectCursor(x, y, width, height, selectType),
		ComponentBasic: *NewComponentBasic(x, y, width, height),
		TextArea:       *textArea,
	}
	return &option
}

func (s *SelectOption) Draw(screen *ebiten.Image) {
	text := s.cursor.GetText() + s.text
	s.TextArea.UpdateText(text)
	s.TextArea.Draw(screen)
}

func (s *SelectOption) Select(selectType SelectType) {
	s.cursor.Select()
}

func (s *SelectOption) Unselect(selectType SelectType) {
	s.cursor.Unselect()
}

func (s *SelectOption) Choose() {
	s.cursor.Choose()
}

func (s *SelectOption) Unchoose() {
	s.cursor.Unchoose()
}

type SelectBox struct {
	ComponentBasic
	image *ebiten.Image

	selectType SelectType     // 多选还是单选
	options    []SelectOption // 选项列表
	selected   util.Set[int]  // 选择的选项，表示选择结果
	choosed    int            // 选中的选项，指示光标位置
	counter    int            // 光标闪烁计数器
}

func NewSelectBox(x, y, width, height int, selectType SelectType) *SelectBox {
	image := ebiten.NewImage(width, height)
	image.Fill(color.Gray{Y: 128})
	selectBox := SelectBox{
		ComponentBasic: *NewComponentBasic(x, y, width, height),
		image:          image,
		selectType:     selectType,
		selected:       util.Set[int]{},
		choosed:        0,
		counter:        0,
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

func (s *SelectBox) GetOptionCount() int {
	return len(s.options)
}

func (s *SelectBox) Draw(screen *ebiten.Image) {
	for _, option := range s.options {
		option.Draw(s.image)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.x), float64(s.y))
	screen.DrawImage(s.image, op)
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

func (s *SelectBox) Choose(index int) {
	s.options[s.choosed].Unchoose()
	s.options[index].Choose()
	s.choosed = index
}

func (s *SelectBox) GetSelected() []int {
	return s.selected.ToSlice()
}

func (s *SelectBox) Update() error {
	s.counter = (s.counter + 1) % 180
	s.options[s.choosed].cursor.UpdateCounter(s.counter)

	if s.selectType == SelectTypeSingle {
		return s.updateSingle()
	} else {
		return s.updateMulti()
	}
}

func (s *SelectBox) updateSingle() error {
	selected := s.GetSelected()
	if len(selected) == 0 {
		return nil
	}
	selectedIndex := selected[0]
	optionCount := s.GetOptionCount()
	if optionCount == 0 {
		return nil
	}

	// 切换选项，同时更新选择结果
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		newIndex := (selectedIndex - 1 + optionCount) % optionCount
		s.Choose(newIndex)
		s.Select(newIndex)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		newIndex := (selectedIndex + 1) % optionCount
		s.Choose(newIndex)
		s.Select(newIndex)
	}

	return nil

}

func (s *SelectBox) updateMulti() error {
	optionCount := s.GetOptionCount()
	if optionCount == 0 {
		return nil
	}

	// 切换选项
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		newIndex := (s.choosed - 1 + optionCount) % optionCount
		s.Choose(newIndex)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		newIndex := (s.choosed + 1) % optionCount
		s.Choose(newIndex)
	}

	// 选择选项
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.Select(s.choosed)
	}

	return nil
}
