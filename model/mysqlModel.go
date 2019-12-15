package model

import (
	"github.com/jinzhu/gorm"
	"log"
	. "proxy_download/initDB"
)

type Mysql struct {
	BaseModel
	Name        string  `gorm:"column:name;not null" binding:"required" json:"name"`
	Ip          string  `gorm:"column:ip;not null" binding:"required" json:"ip"`
	Port        string  `gorm:"column:port;not null" binding:"required" json:"port"`
	Description *string `gorm:"column:description" json:"description"`
	OsUser      string  `gorm:"column:os_user;not null" binding:"required" json:"user"`
	Password    string  `gorm:"column:password;not null" binding:"required" json:"password"`
	DbName      string  `gorm:"column:db_name;not null" binding:"required" json:"db_name"`
	UserId      string  `gorm:"column:user_id;default:'1'" json:"user_id"`
}

var db *gorm.DB

func init() {

	db = DbInit()
}

func (mysql *Mysql) Detail(id int64) (*Mysql, error) {
	err := db.Where("id = ?", id).First(&Mysql{}).Scan(&mysql).Error
	if err != nil {
		return mysql, err
	}
	return mysql, err
}

func Pagination(db *gorm.DB, page, pagesize int) *gorm.DB {
	db = db.Where("user_id = ?", 1).Offset((page - 1) * pagesize).Limit(pagesize)
	return db
}

func (mysql *Mysql) List(page, pagesize int) (mysqls []Mysql, count int, err error) {
	err = Pagination(db.Order("updated_at desc, id desc"), page, pagesize).Find(&mysqls).Error
	if err != nil {
		return
	}
	err = db.Model(&mysql).Count(&count).Error
	if err != nil {
		return
	}
	return
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

func (mysql *Mysql) Delete(id float64) (err error) {
	err = db.Where("id = ?", id).Delete(&mysql).Error
	if err != nil {
		log.Panicln(" delete mysql error", err.Error())
	}
	return
}

func (mysql *Mysql) Deletes(ids []interface{}) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&mysql).Error
	if err != nil {
		log.Panicln("list delete mysql error", err.Error())
		return
	}
	return
}
