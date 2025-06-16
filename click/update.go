// Game.Update逻辑
package click

import (
	"fmt"

	"github.com/Jinvic/Click/click/component"
	"github.com/Jinvic/Click/click/db"
	"github.com/Jinvic/Click/click/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/spf13/cast"
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

	if component.IsComponentJustClicked(g.difficultyArea) {
		g.status = GameStatusDifficultySwitch
		return nil
	}

	return nil
}

func (g *Game) updateRunning() error {
	err := g.gameArea.Update()
	if err != nil {
		return err
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

	// 按下回车键，切换用户
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// 二次确认
		g.status = GameStatusConfirm
		g.confirmArea.SetHintText(fmt.Sprintf("Switch to user: %s ?", g.userSwitchArea.GetUsername()))
		g.confirmArea.SetOnConfirm(func() {
			g.status = GameStatusReady
			newname := g.userSwitchArea.GetUsername()
			g.switchUser(db.GetUser(newname))
			g.status = GameStatusReady
		})
		g.confirmArea.SetOnCancel(func() {
			g.status = GameStatusUserSwitch
		})
		return nil
	}

	// 按下ESC键，返回主界面
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.status = GameStatusReady
		return nil
	}

	err := g.userSwitchArea.Update()
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) updateConfirm() error {
	if g.confirmArea.IsConfirmButtonJustClicked() {
		log.Debug("Confirm")
		g.confirmArea.OnConfirm()
		return nil
	}

	if g.confirmArea.IsCancelButtonJustClicked() {
		log.Debug("Cancel")
		g.confirmArea.OnCancel()
		return nil
	}

	return nil
}

func (g *Game) updateDifficultySwitch() error {
	status := g.difficultySwitchArea.GetStatus()
	switch status {
	case component.DifficultySwitchAreaStatusDifficulty:
		return g.updateDifficultySwitchDifficulty()
	case component.DifficultySwitchAreaStatusValue:
		return g.updateDifficultySwitchValue()
	case component.DifficultySwitchAreaStatusValueInput:
		return g.updateDifficultySwitchValueInput()
	}
	return nil
}

func (g *Game) updateDifficultySwitchDifficulty() error {
	// 按下ESC键，返回主界面
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.status = GameStatusReady
		return nil
	}

	// 按下回车键，切换难度
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		selected := g.difficultySwitchArea.DifficultySelectBox.GetSelected()
		selectedIndex := selected[0]
		var difficulty *component.GameDifficulty
		switch selectedIndex {
		case 0:
			difficulty = &component.GameDifficultyEasy
		case 1:
			difficulty = &component.GameDifficultyMedium
		case 2:
			difficulty = &component.GameDifficultyHard
		case 3: // 自定义难度
			g.difficultySwitchArea.SwitchStatus(component.DifficultySwitchAreaStatusValue)
			return nil
		}

		g.setGameDifficulty(*difficulty)
		g.status = GameStatusReady
		return nil
	}

	err := g.difficultySwitchArea.DifficultySelectBox.Update()
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) updateDifficultySwitchValue() error {
	// 按下ESC键，返回难度选择
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		// 二次确认
		g.status = GameStatusConfirm
		difficulty := g.difficultySwitchArea.GetCustomDifficulty()
		g.confirmArea.SetHintText(fmt.Sprintf("Apply custom difficulty: %v ?", g.difficultySwitchArea.GetCustomDifficulty()))
		g.confirmArea.SetOnConfirm(func() {
			g.setGameDifficulty(difficulty)
			g.status = GameStatusReady
			g.difficultySwitchArea.SwitchStatus(component.DifficultySwitchAreaStatusDifficulty)
		})
		g.confirmArea.SetOnCancel(func() {
			g.status = GameStatusDifficultySwitch
			g.difficultySwitchArea.SwitchStatus(component.DifficultySwitchAreaStatusDifficulty)
		})

		return nil
	}

	// 按下回车键，选择某项数值进行修改
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		value := g.difficultySwitchArea.GetValue()
		log.Debug("old value: ", value)
		g.difficultySwitchArea.ValueInputBox.SetText(cast.ToString(value))
		g.difficultySwitchArea.SwitchStatus(component.DifficultySwitchAreaStatusValueInput)
		return nil
	}

	err := g.difficultySwitchArea.ValueSelectBox.Update()
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) updateDifficultySwitchValueInput() error {
	// 按下ESC键，返回数值项选择
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.difficultySwitchArea.SwitchStatus(component.DifficultySwitchAreaStatusValue)
		return nil
	}

	// 按下回车键，修改数值，返回数值项选择
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.difficultySwitchArea.SetValue()
		log.Debug("new value: ", g.difficultySwitchArea.GetValue())
		g.difficultySwitchArea.SwitchStatus(component.DifficultySwitchAreaStatusValue)
		return nil
	}

	err := g.difficultySwitchArea.ValueInputBox.Update()
	if err != nil {
		return err
	}

	return nil
}
