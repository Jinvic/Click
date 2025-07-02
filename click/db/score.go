package db

import (
	"github.com/Jinvic/Click/click/log"
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	UserID       uint
	DifficultyID uint
	Score        int
}

func GetScore(userID, difficultyID uint) int {
	mutex.Lock()
	defer mutex.Unlock()

	score := Score{}
	err := DB.Model(&Score{}).
		Where("user_id = ? AND difficulty_id = ?", userID, difficultyID).
		First(&score).Error
	if err != nil {
		log.Error(err)
		return 0
	}
	return score.Score
}

func SaveScore(userID, difficultyID uint, score int) {
	mutex.Lock()
	defer mutex.Unlock()

	attr := Score{
		UserID:       userID,
		DifficultyID: difficultyID,
		Score:        score,
	}

	data := Score{}
	err := DB.Model(&Score{}).
		Where("user_id = ? AND difficulty_id = ?", userID, difficultyID).
		Assign(attr).
		FirstOrCreate(&data).Error
	if err != nil {
		log.Error(err)
	}
}
