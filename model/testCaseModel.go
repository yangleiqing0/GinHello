package model

import (
	"log"
)

type TestCase struct {
	BaseModel
	Name                   string `gorm:"column:name;not null" binding:"required" json:"name"`
	Url                    string `gorm:"column:url;not null" binding:"required" json:"url"`
	Data                   string `gorm:"column:data;not null;type:text" binding:"required" json:"data"`
	RegisterVariable       string `gorm:"column:register_variable" json:"register_variable"`
	Regular                string `gorm:"column:regular" json:"regular"`
	Method                 string `gorm:"column:method;not null" binding:"required" json:"method"`
	GroupId                int    `gorm:"column:group_id" json:"group_id"`
	HeaderId               int    `gorm:"column:header_id" json:"header_id"`
	SceneId                int    `gorm:"column:scene_id" json:"scene_id"`
	HopeResult             string `gorm:"column:hope_result;not null" binding:"required" json:"hope_result"`
	IsModel                uint8  `gorm:"column:is_model" json:"is_model"`
	OldSql                 string `gorm:"column:old_sql" json:"old_sql"`
	NewSql                 string `gorm:"column:new_sql" json:"new_sql"`
	OldSqlRegisterVariable string `gorm:"column:old_sql_register_variable" json:"old_sql_register_variable"`
	NewSqlRegisterVariable string `gorm:"column:new_sql_register_variable" json:"new_sql_register_variable"`
	OldSqlHopeResult       string `gorm:"column:old_sql_hope_result" json:"old_sql_hope_result"`
	NewSqlHopeResult       string `gorm:"column:new_sql_hope_result" json:"new_sql_hope_result"`
	OldSqlId               int    `gorm:"column:old_sql_id" json:"old_sql_id"`
	NewSqlId               int    `gorm:"column:new_sql_id" json:"new_sql_id"`
	Description            string `gorm:"column:description"  json:"description"`
	UserId                 int    `gorm:"column:user_id;default:1" json:"user_id"`
	Wait                   Wait   `gorm:"-" json:"wait"`
}

type Wait struct {
	OldWaitSql   string `gorm:"column:old_wait_sql" json:"old_wait_sql"`
	OldWait      string `gorm:"column:old_wait" json:"old_wait"`
	OldWaitTime  string `gorm:"column:old_wait_time" json:"old_wait_time"`
	OldWaitMysql string `gorm:"column:old_wait_mysql" json:"old_wait_mysql"`
	NewWaitSql   string `gorm:"column:new_wait_sql" json:"new_wait_sql"`
	NewWait      string `gorm:"column:new_wait" json:"new_wait"`
	NewWaitTime  string `gorm:"column:new_wait_time" json:"new_wait_time"`
	NewWaitMysql string `gorm:"column:new_wait_mysql" json:"new_wait_mysql"`
}

func (testCase *TestCase) Detail(id int64) (*TestCase, error) {

	err := db.Where("id = ?", id).First(&TestCase{}).Scan(&testCase).Error

	if err != nil {
		return testCase, err
	}
	log.Println("newCase = ", testCase)
	return testCase, err
}

func (testCase *TestCase) List(page, pagesize int64) (testCases []TestCase, count int64, err error) {
	err = Pagination(db.Order("updated_at desc, id desc"), page, pagesize).Find(&testCases).Error
	if err != nil {
		return
	}
	err = db.Model(&testCase).Count(&count).Error
	if err != nil {
		return
	}
	return
}

func (testCase *TestCase) Save() (id int64, err error) {

	err = db.Create(testCase).Error
	if err != nil {
		log.Panicln(" save testCase error", err.Error())
		return
	}
	id = int64(testCase.ID)
	return
}

func (testCase *TestCase) Update() (err error) {

	err = db.Model(&testCase).Update(testCase).Error
	if err != nil {
		log.Panicln(" update testCase error", err.Error())
	}
	return
}

func (testCase *TestCase) Delete(id float64) (err error) {
	err = db.Where("id = ?", id).Delete(&testCase).Error
	if err != nil {
		log.Panicln(" delete testCase error", err.Error())
	}
	return
}

func (testCase *TestCase) Deletes(ids []interface{}) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&testCase).Error
	if err != nil {
		log.Panicln("list delete testCase error", err.Error())
		return
	}
	return
}
