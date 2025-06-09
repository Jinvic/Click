package db

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	MaxScore int
}

func init() {
	// 检查用户表是否为空
	var count int64
	DB.Model(&User{}).Count(&count)

	// 如果没有用户，创建默认用户
	if count == 0 {
		defaultUser := User{
			Username: "Player",
			MaxScore: 0,
		}
		DB.Create(&defaultUser)
	}
}

func SaveUser(user *User) {
	mutex.Lock()
	defer mutex.Unlock()

	// 先查询用户是否存在
	var existingUser User
	result := DB.Where("username = ?", user.Username).First(&existingUser)

	if result.Error == nil {
		// 用户存在，更新记录
		DB.Save(user)
	} else if result.Error == gorm.ErrRecordNotFound {
		// 用户不存在，创建新记录
		DB.Create(user)
	} else {
		fmt.Println(result.Error)
	}
}

func GetUser(username string) *User {
	mutex.Lock()
	defer mutex.Unlock()

	// 先查询用户是否存在
	var user User
	result := DB.Where("username = ?", username).First(&user)

	if result.Error == nil {
		// 用户存在，返回用户
		return &user
	} else if result.Error == gorm.ErrRecordNotFound {
		// 用户不存在，创建该用户
		user := User{
			Username: username,
			MaxScore: 0,
		}
		DB.Create(&user)
		return &user
	} else {
		fmt.Println(result.Error)
		return nil
	}
}
