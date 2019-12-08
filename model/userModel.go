package model

import (
	"github.com/jinzhu/gorm"
	"log"
	."proxy_download/initDB"
)

type User struct {
	gorm.Model
	Name          string `form:"name" gorm:"unique;not null" binding:"required"`
	Email         string `form:"email" binding:"email" gorm:"not null" binding:"required"`
	Password      string `form:"password" gorm:"not null" binding:"required"`
}

func init() {

	db := DbInit()

	if !db.HasTable(&User{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
}

func (user *User) Save() (id int64, err error) {

	err = DbInit().Create(user).Error
	if err != nil {
		log.Panicln("user insert error", err.Error())
		return
	}
	id = int64(user.ID)

	return
}


