package component

import "github.com/hajimehoshi/ebiten/v2"

type GameArea struct {
	x      int
	y      int
	width  int
	height int
}

func NewGameArea(x, y, width, height int) *GameArea {
	// TODO: 实现
	return &GameArea{x: x, y: y, width: width, height: height}
}

func (g *GameArea) Position() (x, y int) {
	return g.x, g.y
}

func (g *GameArea) Draw(screen *ebiten.Image) {
	// TODO: 实现
}
