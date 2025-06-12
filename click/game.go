package click

import (
	"fmt"
	"strings"

	"github.com/Jinvic/Click/click/component"
	"github.com/Jinvic/Click/click/db"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
)

type GameStatus int

const (
	GameStatusReady            GameStatus = iota // 游戏未开始
	GameStatusRunning                            // 游戏进行中
	GameStatusGameOver                           // 游戏结束
	GameStatusUserSwitch                         // 用户选择
	GameStatusDifficultySwitch                   // 难度选择
	GameStatusHelp                               // 帮助界面
	GameStatusConfirm                            // 确认界面
)

const (
	helpText               = "Press Space to start or end the game\nPress R to reset the game\nPress E to exit the game\nPress H to show this help"
	
)

type Game struct {
	status     GameStatus
	clickCount int
	user       *db.User
	difficulty *component.GameDifficulty

	scoreArea      *component.TextArea
	userArea       *component.TextArea
	maxScoreArea   *component.TextArea
	helpArea       *component.MultiTextArea
	difficultyArea *component.DifficultyArea
	gameArea       *component.GameArea
	userSwitchArea *component.UserSwitchArea
	confirmArea    *component.ConfirmArea

	resetButton *component.Button
	exitButton  *component.Button
	startButton *component.Button
	endButton   *component.Button
	helpButton  *component.Button

	components map[GameStatus][]component.Component
}

func NewGame() *Game {
	var user = db.GetUser("Player")
	var difficulty = component.DefaultDifficulty

	var userArea = component.NewTextArea(SCREEN_WIDTH-220, 0, 220, 50, fmt.Sprintf("User: %s", user.Username))
	var scoreArea = component.NewTextArea(SCREEN_WIDTH-220, 60, 220, 50, "Score: 0")
	var maxScoreArea = component.NewTextArea(SCREEN_WIDTH-220, 120, 220, 50, fmt.Sprintf("Max Score: %d", user.MaxScore))
	var gameArea = component.NewGameArea(0, 0, 400, SCREEN_HEIGHT, difficulty)
	var difficultyArea = component.NewDifficultyArea(SCREEN_WIDTH-220, 180, 220, 180, difficulty)

	var helpArea = component.NewMultiTextArea(0, SCREEN_HEIGHT/4, SCREEN_WIDTH, SCREEN_HEIGHT/2, strings.Split(helpText, "\n"))
	var userSwitchArea = component.NewUserSwitchArea(0, SCREEN_HEIGHT/4, SCREEN_WIDTH, SCREEN_HEIGHT/2, user.Username)
	var confirmArea = component.NewConfirmArea(0, SCREEN_HEIGHT/4, SCREEN_WIDTH, SCREEN_HEIGHT/2, "ConfirmArea")

	var resetButton = component.NewButton(
		SCREEN_WIDTH-component.BUTTON_WIDTH-component.BUTTON_WIDTH-20, // 和退出按钮左右间隔20px
		SCREEN_HEIGHT-component.BUTTON_HEIGHT,
		component.BUTTON_WIDTH,
		component.BUTTON_HEIGHT,
		"Reset")
	var exitButton = component.NewButton(
		SCREEN_WIDTH-component.BUTTON_WIDTH,
		SCREEN_HEIGHT-component.BUTTON_HEIGHT,
		component.BUTTON_WIDTH,
		component.BUTTON_HEIGHT,
		"Exit")
	var startButton = component.NewButton(
		SCREEN_WIDTH-component.BUTTON_WIDTH*2-20,
		SCREEN_HEIGHT-component.BUTTON_HEIGHT*2-10,
		component.BUTTON_WIDTH,
		component.BUTTON_HEIGHT,
		"Start")
	var endButton = component.NewButton(
		SCREEN_WIDTH-component.BUTTON_WIDTH*2-20,
		SCREEN_HEIGHT-component.BUTTON_HEIGHT*2-10,
		component.BUTTON_WIDTH*2+20,
		component.BUTTON_HEIGHT,
		"End")
	var helpButton = component.NewButton(
		SCREEN_WIDTH-component.BUTTON_WIDTH,
		SCREEN_HEIGHT-component.BUTTON_HEIGHT*2-10,
		component.BUTTON_WIDTH,
		component.BUTTON_HEIGHT,
		"Help")

	var components = make(map[GameStatus][]component.Component)
	components[GameStatusReady] = []component.Component{
		scoreArea,
		gameArea,
		userArea,
		maxScoreArea,
		difficultyArea,
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
		difficultyArea,
		resetButton,
		exitButton,
		endButton,
	}
	components[GameStatusHelp] = []component.Component{
		helpArea,
		helpButton,
	}
	components[GameStatusUserSwitch] = []component.Component{
		userSwitchArea,
	}
	components[GameStatusConfirm] = []component.Component{
		confirmArea,
	}

	return &Game{
		status:         GameStatusReady,
		user:           user,
		difficulty:     &difficulty,
		scoreArea:      scoreArea,
		userArea:       userArea,
		maxScoreArea:   maxScoreArea,
		helpArea:       helpArea,
		difficultyArea: difficultyArea,
		gameArea:       gameArea,
		userSwitchArea: userSwitchArea,
		confirmArea:    confirmArea,

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
		return g.updateReady()
	case GameStatusRunning:
		return g.updateRunning()
	case GameStatusHelp:
		return g.updateHelp()
	case GameStatusUserSwitch:
		return g.updateUserSwitch()
	case GameStatusConfirm:
		return g.updateConfirm()
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
