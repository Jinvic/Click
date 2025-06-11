package click

import (
	"fmt"

	"github.com/Jinvic/Click/click/component"
	"github.com/Jinvic/Click/click/db"
	"github.com/Jinvic/Click/click/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) updateReady() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) || component.IsComponentJustClicked(g.resetButton) {
		return g.resetGame()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) || component.IsComponentJustClicked(g.exitButton) {
		return g.exitGame()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || component.IsComponentJustClicked(g.startButton) {
		return g.startGame()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) || component.IsComponentJustClicked(g.helpButton) {
		return g.showHelp()
	}
	return nil
}

func (g *Game) updateRunning() error {
	if component.IsComponentJustClicked(g.gameArea) {
		if g.gameArea.IsGameTargetJustClicked() {
			log.Info("Click target")
			g.updateClickCount(g.clickCount + 1)
		} else {
			log.Info("Miss target")
			return g.endGame()
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) || component.IsComponentJustClicked(g.resetButton) {
		return g.resetGame()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) || component.IsComponentJustClicked(g.exitButton) {
		return g.exitGame()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || component.IsComponentJustClicked(g.endButton) {
		return g.endGame()
	}

	return nil
}

func (g *Game) updateHelp() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyH) || component.IsComponentJustClicked(g.helpButton) {
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
	g.gameArea.NewTarget()
	g.gameArea.ShowTarget = true
	return nil
}

func (g *Game) endGame() error {
	log.Info("End game")
	g.status = GameStatusReady
	g.gameArea.ShowTarget = false
	g.gameArea.Clear()
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
