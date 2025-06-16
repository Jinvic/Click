package component

import (
	"fmt"
	"image/color"
	"math"
	"math/rand/v2"
	"time"

	"github.com/Jinvic/Click/click/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameArea struct {
	ComponentBasic
	image *ebiten.Image

	difficulty GameDifficulty
	target     GameTarget
	ShowTarget bool
	timer      *Timer

	onTargetClicked func() error
	onTargetMissed  func() error
}

type GameDifficultyName string

const (
	GameDifficultyNameEasy   GameDifficultyName = "Easy"
	GameDifficultyNameMedium GameDifficultyName = "Medium"
	GameDifficultyNameHard   GameDifficultyName = "Hard"
	GameDifficultyNameCustom GameDifficultyName = "Custom"
)

type GameDifficulty struct {
	Name     GameDifficultyName
	Radius   int
	Speed    int
	Duration int
}

var (
	GameDifficultyEasy = GameDifficulty{
		Name:     GameDifficultyNameEasy,
		Radius:   36,
		Speed:    4,
		Duration: 3000,
	}
	GameDifficultyMedium = GameDifficulty{
		Name:     GameDifficultyNameMedium,
		Radius:   24,
		Speed:    6,
		Duration: 2000,
	}
	GameDifficultyHard = GameDifficulty{
		Name:     GameDifficultyNameHard,
		Radius:   12,
		Speed:    8,
		Duration: 1000,
	}
	DefaultDifficulty = GameDifficultyMedium
)

const (
	TOLERANCE = 10
)

type GameTarget struct {
	x     float64
	y     float64
	angle float64
}

func NewGameArea(x, y, width, height int, difficulty GameDifficulty) *GameArea {
	image := ebiten.NewImage(width, height)
	image.Fill(color.Gray{Y: 128})
	timer := NewTimer(x, y, width, height)
	timer.SetFormat(TimerFormatSecond | TimerFormatMillisecond)
	timer.SetMode(TimerModeCountdown)
	timer.SetLimit(time.Duration(difficulty.Duration) * time.Millisecond)

	gameArea := &GameArea{
		ComponentBasic: *NewComponentBasic(x, y, width, height),
		image:          image,
		difficulty:     difficulty,
		ShowTarget:     false,
		timer:          timer,
	}
	// 设置计时结束回调
	gameArea.timer.SetOnTimerEnd(func() {
		gameArea.onTargetMissed()
	})
	return gameArea
}

func (g *GameArea) Draw(screen *ebiten.Image) {
	if g.ShowTarget {
		g.DrawTarget()
	}
	g.timer.Draw(g.image)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.x), float64(g.y))
	screen.DrawImage(g.image, op)
}

// 随机生成目标
func (g *GameArea) NewTarget() {
	x := rand.Float64()*float64(g.width-2*g.difficulty.Radius) + float64(g.difficulty.Radius)
	y := rand.Float64()*float64(g.height-2*g.difficulty.Radius) + float64(g.difficulty.Radius)
	angle := rand.Float64() * 2 * math.Pi
	log.Debug("NewTarget: " + fmt.Sprintf("x: %f, y: %f, angle: %f", x, y, angle))

	g.target = GameTarget{
		x:     x,
		y:     y,
		angle: angle,
	}
}

// 更新目标位置
func (g *GameArea) UpdateTarget() {
	if !g.ShowTarget {
		return
	}

	g.target.x += float64(g.difficulty.Speed) * math.Cos(g.target.angle)
	g.target.y += float64(g.difficulty.Speed) * math.Sin(g.target.angle)

	// 检查目标是否超出边界
	if g.target.x < float64(g.difficulty.Radius) || g.target.x > float64(g.width-g.difficulty.Radius) {
		g.target.angle = math.Pi - g.target.angle
	}
	if g.target.y < float64(g.difficulty.Radius) || g.target.y > float64(g.height-g.difficulty.Radius) {
		g.target.angle = 2*math.Pi - g.target.angle
	}

	// 确保目标在游戏区域内
	g.target.x = math.Max(float64(g.difficulty.Radius), math.Min(g.target.x, float64(g.width-g.difficulty.Radius)))
	g.target.y = math.Max(float64(g.difficulty.Radius), math.Min(g.target.y, float64(g.height-g.difficulty.Radius)))
}

// 检查目标是否被点击
func (g *GameArea) IsGameTargetJustClicked() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		dx := float64(mx-g.x) - g.target.x
		dy := float64(my-g.y) - g.target.y
		distance := math.Sqrt(dx*dx + dy*dy)
		log.Debug("GameTargetJustClicked: " + fmt.Sprintf("distance: %f, radius: %d, tolerance: %d", distance, g.difficulty.Radius, TOLERANCE))
		if distance <= float64(g.difficulty.Radius+TOLERANCE) {
			return true
		}
	}
	return false
}

// 在游戏区域上绘制目标
func (g *GameArea) DrawTarget() {
	g.image.Fill(color.Gray{Y: 128})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.target.x, g.target.y)
	vector.DrawFilledCircle(g.image, float32(g.target.x), float32(g.target.y), float32(g.difficulty.Radius), color.RGBA{R: 255, G: 0, B: 0, A: 255}, true)
}

// 清空游戏区域
func (g *GameArea) Clear() {
	g.image.Fill(color.Gray{Y: 128})
}

// 设置游戏难度
func (g *GameArea) SetDifficulty(difficulty GameDifficulty) {
	g.difficulty = difficulty
}

func (g *GameArea) SetOnTargetClicked(onTargetClicked func() error) {
	g.onTargetClicked = func() error {
		g.timer.Reset()
		g.timer.Start()
		return onTargetClicked()
	}
}

func (g *GameArea) SetOnTargetMissed(onTargetMissed func() error) {
	g.onTargetMissed = onTargetMissed
}

func (g *GameArea) Update() error {
	if IsComponentJustClicked(g) {
		if g.IsGameTargetJustClicked() {
			g.onTargetClicked()
		} else {
			g.onTargetMissed()
		}
	}

	g.UpdateTarget()
	g.timer.Update()
	return nil
}

func (g *GameArea) StartGame() {
	g.NewTarget()
	g.ShowTarget = true
	g.timer.Start()
}

func (g *GameArea) EndGame() {
	g.ShowTarget = false
	g.Clear()
}

func (g *GameArea) ResetGame() {
	g.image.Fill(color.Gray{Y: 128})
	g.timer.Reset()
}
