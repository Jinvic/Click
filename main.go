package main

import (
	"github.com/Jinvic/Click/click"
	"github.com/Jinvic/Click/click/log"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 320
	windowHeight = 240
	windowTitle  = "Click"
)

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle(windowTitle)
	log.SetLevel(log.LevelInfo)
	ebiten.RunGame(click.NewGame())
}
