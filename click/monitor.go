package click

// Monitor of each operation

import (
	"github.com/Jinvic/Click/click/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Reset the game
func (g *Game) resetGameMonitor() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		log.Info("Reset key pressed")
		g.updateClickCount(0)
	}
	if g.resetButton.IsButtonJustPressed() {
		log.Info("Reset button pressed")
		g.updateClickCount(0)
	}
	return nil
}

// Exit the game
func (g *Game) exitGameMonitor() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		log.Info("Exit key pressed")
		return ebiten.Termination
	}
	if g.exitButton.IsButtonJustPressed() {
		log.Info("Exit button pressed")
		return ebiten.Termination
	}
	return nil
}

// Start the game
func (g *Game) startGameMonitor() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		log.Info("Start key pressed")
		g.updateClickCount(0)
	}
	if g.startButton.IsButtonJustPressed() {
		log.Info("Start button pressed")
		g.updateClickCount(0)
		g.status = GameStatusRunning
	}
	return nil
}

// End the game
func (g *Game) endGameMonitor() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		log.Info("End key pressed")
		g.status = GameStatusReady
	}
	if g.endButton.IsButtonJustPressed() {
		log.Info("End button pressed")
		g.status = GameStatusReady
	}
	return nil
}

// How to play
func (g *Game) helpMonitor() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		log.Info("Help key pressed")
		g.status = GameStatusHelp
	}
	if g.helpButton.IsButtonJustPressed() {
		log.Info("Help button pressed")
		g.status = GameStatusHelp
	}
	return nil
}
