package component

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	Position() (x, y int)
	Draw(screen *ebiten.Image)
}
