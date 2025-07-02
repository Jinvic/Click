package db

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/Jinvic/Click/click/log"
	"gorm.io/gorm"
)

type Difficulty struct {
	gorm.Model
	Name     string
	Radius   int
	Speed    int
	Duration int

	Hash string `gorm:"uniqueIndex"`
}

func (d *Difficulty) GetHash() string {
	hashString := fmt.Sprintf("%s %d %d %d", d.Name, d.Radius, d.Speed, d.Duration)
	hash := sha256.Sum256([]byte(hashString))
	return hex.EncodeToString(hash[:])
}

var (
	DifficultyEasy = Difficulty{
		Name:     "Easy",
		Radius:   24,
		Speed:    4,
		Duration: 2500,
	}
	DifficultyMedium = Difficulty{
		Name:     "Medium",
		Radius:   18,
		Speed:    6,
		Duration: 2000,
	}
	DifficultyHard = Difficulty{
		Name:     "Hard",
		Radius:   12,
		Speed:    8,
		Duration: 1500,
	}
)

func init() {
	// 检查难度表是否为空
	var count int64
	DB.Model(&Difficulty{}).Count(&count)

	// 如果没有难度，创建默认难度
	if count == 0 {
		DifficultyEasy.Hash = DifficultyEasy.GetHash()
		DifficultyMedium.Hash = DifficultyMedium.GetHash()
		DifficultyHard.Hash = DifficultyHard.GetHash()
		DB.Create(&DifficultyEasy)
		DB.Create(&DifficultyMedium)
		DB.Create(&DifficultyHard)	
	}
}

func GetDifficultyId(d *Difficulty) (id uint) {
	mutex.Lock()
	defer mutex.Unlock()

	err := DB.Model(&Difficulty{}).
		Where("hash = ?", d.Hash).
		Assign(d).
		FirstOrCreate(&d).Error
	if err != nil {
		log.Error(err)
		return 0
	}
	return d.ID
}
