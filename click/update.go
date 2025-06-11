// Game.Update逻辑
package click

import (
	"github.com/Jinvic/Click/click/component"
	"github.com/Jinvic/Click/click/db"
	"github.com/Jinvic/Click/click/log"
	"github.com/Jinvic/Click/click/util"
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

	if component.IsComponentJustClicked(g.userArea) {
		g.status = GameStatusUserSwitch
		return nil
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

func (g *Game) updateUserSwitch() error {
	// 输入用户名
	runes := ebiten.AppendInputChars(nil)
	if len(runes) > 0 {
		username := g.userSwitchArea.GetUsername()
		if len(username) < 10 {
			g.userSwitchArea.SetUsername(username + string(runes[0]))
		}
	}

	// 按下退格键，删除字符
	if util.IsKeyLongPressed(ebiten.KeyBackspace) {
		username := g.userSwitchArea.GetUsername()
		if len(username) > 0 {
			g.userSwitchArea.SetUsername(username[:len(username)-1])
		}
	}

	// 按下回车键，切换用户
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.status = GameStatusReady
		newname := g.userSwitchArea.GetUsername()
		g.switchUser(db.GetUser(newname))
		return nil
	}

	// 按下ESC键，返回主界面
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.status = GameStatusReady
		return nil
	}

	g.userSwitchArea.UpdateCursorCounter()
	return nil
}
