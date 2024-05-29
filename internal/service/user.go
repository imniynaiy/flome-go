package service

import (
	"example.com/app-structure/internal/database"
	"example.com/app-structure/internal/model"
	"example.com/app-structure/internal/util"
	"gorm.io/gorm"
)

func getUserRepo() *gorm.DB {
	return database.GlobalDB.Model(&model.User{})
}

func Login(user *model.User) (token string, err error) {
	var userInDb model.User
	err = getUserRepo().Where("username = ?", user.Username).First(&userInDb).Error
	if err != nil {
		return "", err
	}
	err = util.CompareHashAndPassword([]byte(userInDb.Password), user.Password)
	if err != nil {
		return "", err
	}
	return util.CreateJWT(userInDb.Username)
}
