package component

import (
	"fmt"
	"strings"
)

const (
	difficultyTextTemplate = "Difficulty:%s\nRadius:%d\nSpeed:%d\nDuration:%d"
)

type DifficultyArea struct {
	MultiTextArea
	difficulty GameDifficulty
}

func NewDifficultyArea(x, y, width, height int, difficulty GameDifficulty) *DifficultyArea {
	area := DifficultyArea{difficulty: difficulty}
	difficultyText := fmt.Sprintf(difficultyTextTemplate, difficulty.Name, difficulty.Radius, difficulty.Speed, difficulty.Duration)
	area.MultiTextArea = *NewMultiTextArea(x, y, width, height, strings.Split(difficultyText, "\n"))
	return &area
}

func (d *DifficultyArea) SetDifficulty(difficulty GameDifficulty) {
	d.difficulty = difficulty
	difficultyText := fmt.Sprintf(difficultyTextTemplate, difficulty.Name, difficulty.Radius, difficulty.Speed, difficulty.Duration)
	d.UpdateTexts(strings.Split(difficultyText, "\n"))
}
