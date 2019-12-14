package model

import (
	"log"
)

type Variable struct {
	BaseModel
	Name        string  `gorm:"column:name;not null" binding:"required" json:"name"`
	Value       string  `gorm:"column:value;not null;type:text" binding:"required" json:"value"`
	IsPrivate   *uint8  `gorm:"column:is_private" json:"is_private"`
	Description *string `gorm:"column:description"  json:"description"`
	UserId      int     `gorm:"column:user_id;default:1" json:"user_id"`
}

func (variable *Variable) Detail(id int64) (*Variable, error) {
	err := db.Where("id = ?", id).First(&Variable{}).Scan(&variable).Error
	if err != nil {
		return variable, err
	}
	return variable, err
}

func (variable *Variable) List(page, pagesize int) (variables []Variable, count int, err error) {
	err = Pagination(db.Order("updated_at desc, id desc"), page, pagesize).Find(&variables).Error
	if err != nil {
		return
	}
	err = db.Model(&variable).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (variable *Variable) Save() (id int64, err error) {

	err = db.Create(variable).Error
	if err != nil {
		log.Panicln(" save variable error", err.Error())
		return
	}
	id = int64(variable.ID)
	return
}

func (variable *Variable) Update() (err error) {

	err = db.Model(&variable).Update(variable).Error
	if err != nil {
		log.Panicln(" update variable error", err.Error())
	}
	return
}

func (variable *Variable) Delete(id float64) (err error) {
	err = db.Where("id = ?", id).Delete(&variable).Error
	if err != nil {
		log.Panicln(" delete variable error", err.Error())
	}
	return
}

func (variable *Variable) Deletes(ids []interface{}) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&variable).Error
	if err != nil {
		log.Panicln("list delete variable error", err.Error())
		return
	}
	return
}
