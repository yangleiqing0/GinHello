package model

import (
	"github.com/jinzhu/gorm"
	"log"
)

type TestCase struct {
	BaseModel
	Name                   string  `gorm:"column:name;not null" binding:"required" json:"name"`
	Url                    string  `gorm:"column:url;not null" binding:"required" json:"url"`
	Data                   *string `gorm:"column:data;type:text" json:"data"`
	RegisterVariable       *string `gorm:"column:register_variable" json:"register_variable"`
	Regular                *string `gorm:"column:regular" json:"regular"`
	Method                 string  `gorm:"column:method;not null" binding:"required" json:"method"`
	GroupId                *int    `gorm:"column:group_id" json:"group_id"`
	HeaderId               *int    `gorm:"column:header_id" json:"header_id"`
	SceneId                *int    `gorm:"column:scene_id" json:"scene_id"`
	HopeResult             string  `gorm:"column:hope_result;not null" binding:"required" json:"hope_result"`
	IsModel                *int    `gorm:"column:is_model" json:"is_model"`
	OldSql                 *string `gorm:"column:old_sql" json:"old_sql"`
	NewSql                 *string `gorm:"column:new_sql" json:"new_sql"`
	OldSqlRegisterVariable *string `gorm:"column:old_sql_register_variable" json:"old_sql_register_variable"`
	NewSqlRegisterVariable *string `gorm:"column:new_sql_register_variable" json:"new_sql_register_variable"`
	OldSqlHopeResult       *string `gorm:"column:old_sql_hope_result" json:"old_sql_hope_result"`
	NewSqlHopeResult       *string `gorm:"column:new_sql_hope_result" json:"new_sql_hope_result"`
	OldSqlId               *int    `gorm:"column:old_sql_id" json:"old_sql_id"`
	NewSqlId               *int    `gorm:"column:new_sql_id" json:"new_sql_id"`
	Description            *string `gorm:"column:description"  json:"description"`
	UserId                 int     `gorm:"column:user_id;default:1" json:"user_id"`
	Wait                   Wait    `gorm:"-" json:"wait"`
}

func (testCase *TestCase) Detail(id int) (*TestCase, error) {

	err := db.Where("id = ?", id).First(&TestCase{}).Scan(&testCase).Error

	if err != nil {
		return testCase, err
	}

	err1 := db.Table("waits").Where("testcase_id = ?", id).Find(&testCase.Wait).Error
	if err1 != nil && err1 != gorm.ErrRecordNotFound {
		return testCase, err1
	}

	log.Println("newCase = ", testCase)
	return testCase, nil
}

func (testCase *TestCase) List(page, pagesize int) (testCases []TestCase, count int, err error) {
	err = Pagination(db.Where("scene_id = ?", 0).Order("updated_at desc, id desc"), page, pagesize).Find(&testCases).Error
	if err != nil {
		return
	}
	err = db.Model(&testCase).Where("scene_id = ?", 0).Count(&count).Error
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

	var wait = &testCase.Wait
	wait.TestCaseId = &testCase.ID

	err1 := wait.Save()
	if err1 != nil {
		log.Panicln(" save testCase.wait error", err.Error())
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

	var wait = &testCase.Wait

	err1 := wait.Update(testCase.ID)
	if err1 != nil {
		log.Panicln(" update testCase.wait error", err.Error())
		return
	}
	return
}

func (testCase *TestCase) Delete(id float64) (err error) {
	err = db.Where("id = ?", id).Delete(&testCase).Error
	if err != nil {
		log.Panicln(" delete testCase error", err.Error())
		return
	}
	_ = db.Table("waits").Where("testcase_id = ?", id).Delete(testCase.Wait).Error
	return
}

func (testCase *TestCase) Deletes(ids []interface{}) (err error) {
	err = db.Where("id IN (?)", ids).Delete(&testCase).Error
	if err != nil {
		log.Panicln("list delete testCase error", err.Error())
		return
	}
	_ = db.Table("waits").Where("testcase_id IN (?)", ids).Delete(testCase.Wait).Error
	return
}
