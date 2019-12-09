package model

import (
	"github.com/jinzhu/gorm"
	"log"
	. "proxy_download/initDB"
)

type Mysql struct {
	gorm.Model
	Ip          string `form:"ip" gorm:"column:ip;not null" binding:"required" json:"ip"`
	Port        int    `form:"port" gorm:"not null" binding:"required,min=1,max=65535" json:"port"`
	Description string `form:"description" json:"description"`
	UserId      string `form:"user_id" gorm:"not null" binding:"required" json:"user_id"`
	OsUser      string `form:"os_user" gorm:"not null" binding:"required" json:"os_user"`
	Password    string `form:"password" gorm:"not null" binding:"required" json:"password"`
	DbName      string `form:"db_name" gorm:"not null" binding:"required" json:"db_name"`
}

var db *gorm.DB

func init() {

	db = DbInit()

	if !db.HasTable(&Mysql{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Mysql{}).Error; err != nil {
			panic(err)
		}
	}
}

func (mysql *Mysql) Save() (id int64, err error) {

	err = db.Create(mysql).Error
	if err != nil {
		log.Panicln(" save mysql error", err.Error())
		return
	}
	id = int64(mysql.ID)
	return
}
