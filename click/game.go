package click

import (
	"fmt"
	"strings"

	"github.com/Jinvic/Click/click/component"
	"github.com/Jinvic/Click/click/db"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 320
	SCREEN_HEIGHT = 240
	BUTTON_WIDTH  = 50
	BUTTON_HEIGHT = 20
)

type GameStatus int

const (
	GameStatusReady            GameStatus = iota // 游戏未开始
	GameStatusRunning                            // 游戏进行中
	GameStatusGameOver                           // 游戏结束
	GameStatusUserSwitch                         // 用户选择
	GameStatusDifficultySwitch                   // 难度选择
	GameStatusHelp                               // 帮助界面
)

const (
	helpText = "Press Space to start or end the game\nPress R to reset the game\nPress E to exit the game\nPress H to show this help"
)

type Game struct {
	status     GameStatus
	clickCount int
	user       *db.User

	scoreArea    *component.TextArea
	userArea     *component.TextArea
	maxScoreArea *component.TextArea
	helpArea     *component.MultiTextArea
	gameArea     *component.GameArea

	resetButton *component.Button
	exitButton  *component.Button
	startButton *component.Button
	endButton   *component.Button
	helpButton  *component.Button

	components map[GameStatus][]component.Component
}

func NewGame() *Game {
	var user = db.GetUser("Player")

	var scoreArea = component.NewTextArea(0, 0, 120, 20, "Score: 0")
	var gameArea = component.NewGameArea(0, 30, SCREEN_WIDTH, SCREEN_HEIGHT-BUTTON_HEIGHT-40) // 和其他组件上下间隔10px
	var userArea = component.NewTextArea(SCREEN_WIDTH-120, 0, 120, 20, fmt.Sprintf("User: %s", user.Username))
	var maxScoreArea = component.NewTextArea(0, SCREEN_HEIGHT-BUTTON_HEIGHT, 120, 20, fmt.Sprintf("Max Score: %d", user.MaxScore))
	var helpArea = component.NewMultiTextArea(0, SCREEN_HEIGHT/4, SCREEN_WIDTH, SCREEN_HEIGHT/2, strings.Split(helpText, "\n"))

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
	var startButton = component.NewButton(
		(SCREEN_WIDTH-BUTTON_WIDTH)/2, // 居中
		SCREEN_HEIGHT-BUTTON_HEIGHT,
		BUTTON_WIDTH,
		BUTTON_HEIGHT,
		"Start")
	var endButton = component.NewButton(
		(SCREEN_WIDTH-BUTTON_WIDTH)/2, // 居中
		SCREEN_HEIGHT-BUTTON_HEIGHT,
		BUTTON_WIDTH,
		BUTTON_HEIGHT,
		"End")
	var helpButton = component.NewButton(
		(SCREEN_WIDTH-BUTTON_WIDTH)/2, // 居中
		0,
		BUTTON_WIDTH,
		BUTTON_HEIGHT,
		"Help")

	var components = make(map[GameStatus][]component.Component)
	components[GameStatusReady] = []component.Component{
		scoreArea,
		gameArea,
		userArea,
		maxScoreArea,
		resetButton,
		exitButton,
		startButton,
		helpButton,
	}
	components[GameStatusRunning] = []component.Component{
		scoreArea,
		gameArea,
		userArea,
		maxScoreArea,
		resetButton,
		exitButton,
		endButton,
	}
	components[GameStatusHelp] = []component.Component{
		helpArea,
		helpButton,
	}

	return &Game{
		status: GameStatusReady,
		user:   user,

		scoreArea:    scoreArea,
		gameArea:     gameArea,
		userArea:     userArea,
		maxScoreArea: maxScoreArea,
		helpArea:     helpArea,

		resetButton: resetButton,
		exitButton:  exitButton,
		startButton: startButton,
		endButton:   endButton,
		helpButton:  helpButton,

		components: components,
	}
}

func (g *Game) Update() error {
	switch g.status {
	case GameStatusReady:
		g.readyLogic()
	case GameStatusRunning:
		g.runningLogic()
	case GameStatusHelp:
		g.helpLogic()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, component := range g.components[g.status] {
		component.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
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
