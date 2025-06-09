package click

import (
	"fmt"

	"github.com/Jinvic/Click/click/db"
	"github.com/Jinvic/Click/click/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) updateReady() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) || g.resetButton.IsButtonJustPressed() {
		return g.resetGame()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) || g.exitButton.IsButtonJustPressed() {
		return g.exitGame()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || g.startButton.IsButtonJustPressed() {
		return g.startGame()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) || g.helpButton.IsButtonJustPressed() {
		return g.showHelp()
	}
	return nil
}

func (g *Game) updateRunning() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.updateClickCount(g.clickCount + 1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) || g.resetButton.IsButtonJustPressed() {
		return g.resetGame()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) || g.exitButton.IsButtonJustPressed() {
		return g.exitGame()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || g.endButton.IsButtonJustPressed() {
		return g.endGame()
	}

	return nil
}

func (g *Game) updateHelp() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyH) || g.helpButton.IsButtonJustPressed() {
		return g.closeHelp()
	}
	return nil
}

func (g *Game) updateClickCount(c int) {
	g.clickCount = c
	g.scoreArea.UpdateText(fmt.Sprintf("Score: %d", g.clickCount))
	if g.clickCount > g.user.MaxScore {
		g.user.MaxScore = g.clickCount
		db.SaveUser(g.user)
		g.maxScoreArea.UpdateText(fmt.Sprintf("Max Score: %d", g.user.MaxScore))
	}
}

func (g *Game) resetGame() error {
	log.Info("Reset game")
	g.updateClickCount(0)
	return nil
}

func (g *Game) exitGame() error {
	log.Info("Exit game")
	return ebiten.Termination
}

func (g *Game) startGame() error {
	log.Info("Start game")
	g.updateClickCount(0)
	g.status = GameStatusRunning
	return nil
}

func (g *Game) endGame() error {
	log.Info("End game")
	g.status = GameStatusReady
	return nil
}

func (g *Game) showHelp() error {
	log.Info("Show help")
	g.status = GameStatusHelp
	return nil
}

func (g *Game) closeHelp() error {
	log.Info("Close help")
	g.status = GameStatusReady
	return nil
}
