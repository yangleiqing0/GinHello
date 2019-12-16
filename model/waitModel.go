package model

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Wait struct {
	BaseModel
	OldWaitSql   *string `gorm:"column:old_wait_sql" json:"old_wait_sql"`
	OldWait      *string `gorm:"column:old_wait" json:"old_wait"`
	OldWaitTime  *string `gorm:"column:old_wait_time" json:"old_wait_time"`
	OldWaitMysql *int    `gorm:"column:old_wait_mysql" json:"old_wait_mysql"`
	NewWaitSql   *string `gorm:"column:new_wait_sql" json:"new_wait_sql"`
	NewWait      *string `gorm:"column:new_wait" json:"new_wait"`
	NewWaitTime  *string `gorm:"column:new_wait_time" json:"new_wait_time"`
	NewWaitMysql *int    `gorm:"column:new_wait_mysql" json:"new_wait_mysql"`
	TestCaseId   *int    `gorm:"column:testcase_id" sql:"index" json:"testcase_id"`
}

func (wait *Wait) Save() (err error) {

	err = db.Create(wait).Error
	if err != nil {
		log.Panicln(" save wait error", err.Error())
		return
	}
	return
}

func (wait *Wait) Update(testCaseId int) (err error) {
	var result Wait

	queryWait := db.Table("waits").Where("testcase_id = ?", testCaseId)
	err1 := queryWait.Find(&result).Error

	if err1 == gorm.ErrRecordNotFound {

		err = wait.Save()
		return err
	}

	err = db.Model(&wait).Update(wait).Error
	if err != nil {
		log.Panicln(" update wait error", err.Error())
	}
	return
}
