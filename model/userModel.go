package model

import (
	"log"
)

type User struct {
	BaseModel
	Name     string `form:"name" gorm:"unique;not null" binding:"required"`
	Email    string `form:"email" binding:"email" gorm:"not null" binding:"required"`
	Password string `form:"password" gorm:"not null" binding:"required"`
}

func init() {

}

func (user *User) Save() (id int64, err error) {

	err = db.Create(user).Error
	if err != nil {
		log.Panicln("user insert error", err.Error())
		return
	}
	id = int64(user.ID)
	return
}
