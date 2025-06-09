package db

import "gorm.io/gorm"

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
