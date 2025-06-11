package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Component interface {
	Position() (x, y int)
	Size() (width, height int)
	Draw(screen *ebiten.Image)
}

func IsComponentJustClicked(c Component) bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		x, y := c.Position()
		width, height := c.Size()
		if mx >= x && mx <= x+width && my >= y && my <= y+height {
			return true
		}
	}
	return false
}
