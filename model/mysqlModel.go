package model

import (
	"github.com/jinzhu/gorm"
	"log"
	. "proxy_download/initDB"
)

type Mysql struct {
	BaseModel
	Name        string `gorm:"column:name;not null" binding:"required" json:"name"`
	Ip          string `gorm:"column:ip;not null" binding:"required" json:"ip"`
	Port        string `gorm:"column:port;not null" binding:"required" json:"port"`
	Description string `gorm:"column:description" json:"description"`
	OsUser      string `gorm:"column:os_user;not null" binding:"required" json:"user"`
	Password    string `gorm:"column:password;not null" binding:"required" json:"password"`
	DbName      string `gorm:"column:db_name;not null" binding:"required" json:"db_name"`
	UserId      string `gorm:"column:user_id;default:'1'" json:"user_id"`
}

var db *gorm.DB

func init() {

	db = DbInit()
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

func (mysql *Mysql) Update() (err error) {

	err = db.Model(&mysql).Update(mysql).Error
	if err != nil {
		log.Panicln(" update mysql error", err.Error())
	}
	return
}

func (mysql *Mysql) Delete() (err error) {
	err = db.Delete(&mysql).Error
	if err != nil {
		log.Panicln(" delete mysql error", err.Error())
	}
	return
}
