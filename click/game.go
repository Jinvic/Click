package click

import (
	"fmt"

	"github.com/Jinvic/Click/click/component"
	"github.com/Jinvic/Click/click/db"
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
	clickCount   int
	user         *db.User
	scoreArea    *component.TextArea
	gameArea     *component.GameArea
	userArea     *component.TextArea
	maxScoreArea *component.TextArea
	resetButton  *component.Button
	exitButton   *component.Button
	components   []component.Component
}

func NewGame() *Game {
	var user db.User
	err := db.DB.First(&user, "username = ?", "Player").Error
	if err != nil {
		panic(err)
	}
	var scoreArea = component.NewTextArea(0, 0, 120, 20, "Score: 0")
	var gameArea = component.NewGameArea(0, 30, SCREEN_WIDTH, SCREEN_HEIGHT-BUTTON_HEIGHT-40) // 和其他组件上下间隔10px
	var userArea = component.NewTextArea(SCREEN_WIDTH-120, 0, 120, 20, fmt.Sprintf("User: %s", user.Username))
	var maxScoreArea = component.NewTextArea(0, SCREEN_HEIGHT-BUTTON_HEIGHT, 120, 20, fmt.Sprintf("Max Score: %d", user.MaxScore))
	var resetButton = component.NewButton(
		SCREEN_WIDTH-BUTTON_WIDTH-BUTTON_WIDTH-20, // 和退出按钮左右间隔20px
		SCREEN_HEIGHT-BUTTON_HEIGHT,
		BUTTON_WIDTH,
		BUTTON_HEIGHT,
		"Reset")
	var exitButton = component.NewButton(
		SCREEN_WIDTH-BUTTON_WIDTH,
		SCREEN_HEIGHT-BUTTON_HEIGHT,
		BUTTON_WIDTH,
		BUTTON_HEIGHT,
		"Exit")
	return &Game{
		user:         &user,
		scoreArea:    scoreArea,
		gameArea:     gameArea,
		userArea:     userArea,
		maxScoreArea: maxScoreArea,
		resetButton:  resetButton,
		exitButton:   exitButton,
		components: []component.Component{
			scoreArea,
			gameArea,
			userArea,
			maxScoreArea,
			resetButton,
			exitButton,
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
	if g.clickCount > g.user.MaxScore {
		g.user.MaxScore = g.clickCount
		db.DB.Save(&g.user)
		g.maxScoreArea.UpdateText(fmt.Sprintf("Max Score: %d", g.user.MaxScore))
	}
}
