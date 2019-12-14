package model

import (
	"log"
)

type Header struct {
	BaseModel
	Name        string  `gorm:"column:name;not null" binding:"required" json:"name"`
	Value       string  `gorm:"column:value;not null;type:text" binding:"required" json:"value"`
	Description *string `gorm:"column:description"  json:"description"`
	UserId      int     `gorm:"column:user_id;default:1" json:"user_id"`
}

type HeaderSimple struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (header *Header) Detail(id int64) (*Header, error) {
	err := db.Where("id = ?", id).First(&Header{}).Scan(&header).Error
	if err != nil {
		return header, err
	}
	return header, err
}

func (header *Header) List(page, pagesize int) (headers []Header, count int64, err error) {
	err = Pagination(db.Order("updated_at desc, id desc"), page, pagesize).Find(&headers).Error
	if err != nil {
		return
	}
	err = db.Model(&header).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (header *Header) ListAll() (headers []HeaderSimple, err error) {
	err = db.Table("headers").Select("id, name").Scan(&headers).Error
	if err != nil {
		return
	}
	return
}

func (header *Header) Save() (id int64, err error) {

	err = db.Create(header).Error
	if err != nil {
		log.Panicln(" save header error", err.Error())
		return
	}
	id = int64(header.ID)
	return
}

func (header *Header) Update() (err error) {

	err = db.Model(&header).Update(header).Error
	if err != nil {
		log.Panicln(" update header error", err.Error())
	}
	return
}

func (header *Header) Delete(id float64) (err error) {
	err = db.Where("id = ?", id).Delete(&header).Error
	if err != nil {
		log.Panicln(" delete header error", err.Error())
	}
	return
}

func (header *Header) Deletes(ids []interface{}) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&header).Error
	if err != nil {
		log.Panicln("list delete header error", err.Error())
		return
	}
	return
}
