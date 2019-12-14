package model

import (
	"log"
)

type Group struct {
	BaseModel
	Name        string  `gorm:"column:name;not null" binding:"required" json:"name"`
	Description *string `gorm:"column:description"  json:"description"`
	UserId      int     `gorm:"column:user_id;default:1" json:"user_id"`
}

type GroupSimple struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (group *Group) Detail(id int) (*Group, error) {
	err := db.Where("id = ?", id).First(&Group{}).Scan(&group).Error
	if err != nil {
		return group, err
	}
	return group, err
}

func (group *Group) List(page, pagesize int) (groups []*Group, count int, err error) {
	err = Pagination(db.Order("updated_at desc, id desc"), page, pagesize).Find(&groups).Error
	if err != nil {
		return
	}
	err = db.Model(&group).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (group *Group) ListAll() (groups []*GroupSimple, err error) {
	err = db.Table("groups").Select("id, name").Scan(&groups).Error
	if err != nil {
		return
	}
	return
}

func (group *Group) Save() (id int64, err error) {

	err = db.Create(group).Error
	if err != nil {
		log.Panicln(" save group error", err.Error())
		return
	}
	id = int64(group.ID)
	return
}

func (group *Group) Update() (err error) {

	err = db.Model(&group).Update(group).Error
	if err != nil {
		log.Panicln(" update group error", err.Error())
	}
	return
}

func (group *Group) Delete(id float64) (err error) {
	err = db.Where("id = ?", id).Delete(&group).Error
	if err != nil {
		log.Panicln(" delete group error", err.Error())
	}
	return
}

func (group *Group) Deletes(ids []interface{}) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&group).Error
	if err != nil {
		log.Panicln("list delete group error", err.Error())
		return
	}
	return
}
