package click

// Logic of each game status

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func (g *Game) readyLogic() error {
	g.resetGameMonitor()
	g.exitGameMonitor()
	g.startGameMonitor()
	g.helpMonitor()

	return nil
}

func (g *Game) runningLogic() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.updateClickCount(g.clickCount + 1)
	}
	g.resetGameMonitor()
	g.exitGameMonitor()
	g.endGameMonitor()
	return nil
}

func (g *Game) helpLogic() error {
	g.helpMonitor()
	return nil
}
