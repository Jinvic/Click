package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Component interface {
	Position() (x, y int)         // 相对位置
	AbsolutePosition() (x, y int) // 绝对位置
	Size() (width, height int)
	Draw(screen *ebiten.Image)
}

type ComponentBasic struct {
	father        Component
	x, y          int
	width, height int
}

func (c *ComponentBasic) Position() (x, y int) {
	return c.x, c.y
}

func (c *ComponentBasic) AbsolutePosition() (x, y int) {
	if c.father != nil {
		fx, fy := c.father.AbsolutePosition()
		return fx + c.x, fy + c.y
	}
	return c.x, c.y
}

func (c *ComponentBasic) Size() (width, height int) {
	return c.width, c.height
}

func NewComponentBasic(x, y, width, height int) *ComponentBasic {
	return &ComponentBasic{x: x, y: y, width: width, height: height}
}

func IsComponentJustClicked(c Component) bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		x, y := c.AbsolutePosition()
		width, height := c.Size()
		if mx >= x && mx <= x+width && my >= y && my <= y+height {
			return true
		}
	}
	return false
}
