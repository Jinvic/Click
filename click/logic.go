// 自定义逻辑
package click

import (
	"fmt"

	"github.com/Jinvic/Click/click/component"
	"github.com/Jinvic/Click/click/db"
	"github.com/Jinvic/Click/click/log"
	"github.com/hajimehoshi/ebiten/v2"
)

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

// 切换用户
func (g *Game) switchUser(user *db.User) error {
	log.Info("Switch user:", user.Username)
	g.user = user
	g.userArea.UpdateText(fmt.Sprintf("User: %s", g.user.Username))
	g.scoreArea.UpdateText(fmt.Sprintf("Score: %d", 0))
	g.maxScoreArea.UpdateText(fmt.Sprintf("Max Score: %d", g.user.MaxScore))
	return nil
}

// 设置游戏难度
func (g *Game) setGameDifficulty(difficulty component.GameDifficulty) error {
	log.Info("Set game difficulty:", difficulty)
	g.difficulty = &difficulty
	g.gameArea.SetDifficulty(difficulty)
	g.difficultyArea.SetDifficulty(difficulty)
	return nil
}
