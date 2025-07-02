package db

import (
	"github.com/Jinvic/Click/click/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
}

func init() {
	// 检查用户表是否为空
	var count int64
	DB.Model(&User{}).Count(&count)

	// 如果没有用户，创建默认用户
	if count == 0 {
		defaultUser := User{
			Username: "Player",
		}
		DB.Create(&defaultUser)
	}
}

func SaveUser(user *User) {
	mutex.Lock()
	defer mutex.Unlock()

	err := DB.Where("username = ?", user.Username).
		Assign(user).
		FirstOrCreate(user).Error
	if err != nil {
		log.Error(err)
	}
}

func GetUser(username string) *User {
	mutex.Lock()
	defer mutex.Unlock()

	user := User{
		Username: username,
	}
	err := DB.Where("username = ?", username).
		Assign(user).
		FirstOrCreate(&user).Error
	if err != nil {
		log.Error(err)
		return nil
	}
	return &user
}
