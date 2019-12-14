package model

import (
	"log"
)

type Email struct {
	BaseModel
	Name        string `gorm:"column:name;not null" binding:"required" json:"name"`
	Subject     string `gorm:"column:subject;not null;size:50" binding:"required" json:"subject"`
	ToUserList  string `gorm:"column:to_user_list;not null;" binding:"required" json:"to_user_list"`
	EmailMethod uint8  `gorm:"column:email_method;not null" binding:"required" json:"email_method"`
	UserId      int    `gorm:"column:user_id;default:1" json:"user_id"`
}

func (email *Email) Detail(id int) (*Email, error) {
	err := db.Where("id = ?", id).First(&Email{}).Scan(&email).Error
	if err != nil {
		return email, err
	}
	return email, err
}

func (email *Email) List(page, pagesize int) (emails []Email, count int, err error) {
	err = Pagination(db.Order("updated_at desc, id desc"), page, pagesize).Find(&emails).Error
	if err != nil {
		return
	}
	err = db.Model(&email).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (email *Email) Save() (id int, err error) {

	err = db.Create(email).Error
	if err != nil {
		log.Panicln(" save email error", err.Error())
		return
	}
	id = email.ID
	return
}

func (email *Email) Update() (err error) {

	err = db.Model(&email).Update(email).Error
	if err != nil {
		log.Panicln(" update email error", err.Error())
	}
	return
}

func (email *Email) Delete(id float64) (err error) {
	err = db.Where("id = ?", id).Delete(&email).Error
	if err != nil {
		log.Panicln(" delete email error", err.Error())
	}
	return
}

func (email *Email) Deletes(ids []interface{}) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&email).Error
	if err != nil {
		log.Panicln("list delete email error", err.Error())
		return
	}
	return
}
