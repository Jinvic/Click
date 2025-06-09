package click

import (
	"fmt"

	"github.com/Jinvic/Click/click/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	SCREEN_WIDTH  = 320
	SCREEN_HEIGHT = 240
	BUTTON_WIDTH  = 50
	BUTTON_HEIGHT = 20
)

type Game struct {
	clickCount  int
	scoreArea   *component.TextArea
	resetButton *component.Button
	components  []component.Component
}

func NewGame() *Game {
	var scoreArea = component.NewTextArea(0, 0, 80, 20, "Score: 0")
	var resetButton = component.NewButton(
		SCREEN_WIDTH-BUTTON_WIDTH,
		SCREEN_HEIGHT-BUTTON_HEIGHT,
		BUTTON_WIDTH,
		BUTTON_HEIGHT,
		"Reset")
	return &Game{
		scoreArea:   scoreArea,
		resetButton: resetButton,
		components: []component.Component{
			scoreArea,
			resetButton,
		},
	}
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.UpdateCount(g.clickCount + 1)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.UpdateCount(0)
	}
	if g.resetButton.IsButtonJustPressed() {
		fmt.Println("Reset button pressed")
		g.UpdateCount(0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, component := range g.components {
		component.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) UpdateCount(c int) {
	g.clickCount = c
	g.scoreArea.UpdateText(fmt.Sprintf("Score: %d", g.clickCount))
}
