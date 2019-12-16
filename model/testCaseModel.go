package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"proxy_download/common"
	"strings"
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

func (testCase *TestCase) Save() (id int, err error) {

	err = db.Create(testCase).Error
	if err != nil {
		log.Panicln(" save testCase error", err.Error())
		return
	}

	testCase.SaveCreateVariable()

	var wait = &testCase.Wait
	wait.TestCaseId = &testCase.ID

	err1 := wait.Save()
	if err1 != nil {
		log.Panicln(" save testCase.wait error", err.Error())
		return
	}

	id = testCase.ID
	return
}

func (testCase *TestCase) Update() (err error) {

	err = db.Model(&testCase).Update(testCase).Error
	if err != nil {
		log.Panicln(" update testCase error", err.Error())
	}

	testCase.SaveCreateVariable()
	var wait = &testCase.Wait

	err1 := wait.Update(testCase.ID)
	if err1 != nil {
		log.Panicln(" update testCase.wait error", err.Error())
		return
	}
	return
}

func (testCase *TestCase) Delete(id float64) (err error) {
	testCase.DelCaseDelVariable(id)
	err = db.Where("id = ?", id).Delete(&testCase).Error
	if err != nil {
		log.Panicln(" delete testCase error", err.Error())
		return
	}
	_ = db.Where("testcase_id = ?", id).Delete(&Wait{}).Error
	return
}

func (testCase *TestCase) Deletes(ids []interface{}) (err error) {
	type variable struct {
		OldSqlRegisterVariable string `json:"old_sql_register_variable"`
		NewSqlRegisterVariable string `json:"new_sql_register_variable"`
	}
	var names []variable
	var variableNames []string
	db.Table("test_cases").Select("old_sql_register_variable, new_sql_register_variable").Where("id IN (?)", ids).Scan(&names)
	for i := 0; i < len(names); i++ {
		if names[i].OldSqlRegisterVariable != "" {
			variableNames = append(variableNames, names[i].OldSqlRegisterVariable)
		}
		if names[i].NewSqlRegisterVariable != "" {
			variableNames = append(variableNames, names[i].NewSqlRegisterVariable)
		}
	}

	err = db.Where("id IN (?)", ids).Delete(&testCase).Error
	db.Where("name IN (?)", variableNames).Delete(&Variable{})
	if err != nil {
		log.Panicln("list delete testCase error", err.Error())
		return
	}
	db.Where("testcase_id IN (?)", ids).Delete(&Wait{})
	return
}

func (testCase *TestCase) SaveCreateVariable() {
	if testCase.OldSqlRegisterVariable != nil {
		if len(strings.Trim(*testCase.OldSqlRegisterVariable, " ")) > 0 {
			var variable Variable
			variable.Name = *testCase.OldSqlRegisterVariable
			_, _ = variable.Save()
		}
	}
	if testCase.NewSqlRegisterVariable != nil {
		if len(strings.Trim(*testCase.NewSqlRegisterVariable, " ")) > 0 {
			var variable Variable
			variable.Name = *testCase.NewSqlRegisterVariable
			_, _ = variable.Save()
		}
	}

}

func (testCase *TestCase) DelCaseDelVariable(id float64) {
	var testCase2 TestCase
	db.Where("id = ?", id).Find(&testCase2)
	log.Println("DelCaseDelVariable ", testCase2.OldSqlRegisterVariable, testCase2.NewSqlRegisterVariable)
	if testCase2.OldSqlRegisterVariable != nil {
		db.Where("name = ?", testCase2.OldSqlRegisterVariable).Delete(&Variable{})
	}
	if testCase2.NewSqlRegisterVariable != nil {
		db.Where("name = ?", testCase2.NewSqlRegisterVariable).Delete(&Variable{})
	}
}

func UpdateCaseRegisterNameValidate(testCaseId, userId int, registerVariable string) (result bool, err error) {
	registerVariables := strings.Split(registerVariable, ",")

	var testCase = struct {
		RegisterVariable string `json:"register_variable"`
	}{}

	err = db.Table("test_cases").Select("register_variable").Where("id = ?", testCaseId).Scan(&testCase).Error
	if err != nil {
		fmt.Println("UpdateCaseRegisterNameValidate query register_variable err = ", err)
		return
	}
	caseRegisterVariable := testCase.RegisterVariable
	caseRegisterVariables := strings.Split(caseRegisterVariable, ",")
	if len(common.SliceToMap(registerVariables)) != len(registerVariables) {
		return
	}

	if SliceEq(registerVariables, caseRegisterVariables) {
		return true, nil
	}

	for _, vbName := range registerVariables {
		variable, err := QueryVariable(vbName, userId)
		var count int
		if err != nil {
			count, _ = QueryVariableCount(vbName, 0, userId)
		} else {
			count, _ = QueryVariableCount(vbName, variable.ID, userId)
		}

		if count != 0 {
			return false, nil
		}
	}
	return true, nil
}

func RegularValidate(regular string) (result bool) {

	regulars := strings.Split(regular, ",")
	for _, re := range regulars {
		if strings.Index(re, "$") != -1 {
			if re[1] != '.' || re[0] != '$' {
				return false
			}

			res := strings.Split(re, ".")
			if len(res) != len(common.SliceToMap(res)) {
				return false
			} else {
				if strings.Index(re, " ") != -1 {
					return false
				}
			}
		}
	}

	return true
}

func HopeResultValidate(hopeResult string) bool {

	hopeResults := strings.Split(hopeResult, ",")
	for _, hope := range hopeResults {
		hopes := strings.Split(hope, ":")
		if len(hopes) == 2 {
			r, _ := common.IsStringSliceHas(hopes[0], []string{"包含", "不包含", "等于", "不等于"})
			if r == false {
				return false
			}
			if hopes[1] == "" {
				return false
			}
		}
		if len(hopes) == 1 {
			return false
		}
	}
	return true
}

func CreateSqlVariableValidate(SqlVariable string, UserId int) (result bool, err error) {
	var count int
	if SqlVariable == "" {
		return true, nil
	}
	err = db.Table("variables").Where("name = ? and user_id = ?", SqlVariable, UserId).Count(&count).Error
	if err != nil {
		fmt.Println("CreateSqlVariableValidate count err = ", err)
		return false, err
	}

	return count == 0, nil
}

func UpdateSqlVariableValidate(SqlVariable string, Id, UserId int) (result bool, err error) {
	var params = struct {
		OldSqlRegisterVariable string `json:"old_sql_register_variable"`
		NewSqlRegisterVariable string `json:"new_sql_register_variable"`
	}{}
	if SqlVariable == "" {
		return true, nil
	}

	err = db.Table("test_cases").Where("id = ?", Id).Scan(&params).Error
	if err != nil {
		fmt.Println("UpdateSqlVariableValidate query test_ cases err = ", err)
		return
	}

	r, _ := common.IsStringSliceHas(SqlVariable, []string{params.OldSqlRegisterVariable, params.NewSqlRegisterVariable})
	if r {
		return true, nil
	}

	type variableSimple struct {
		Id int `json:"id"`
	}

	var variable variableSimple

	err = db.Table("variables").Where("name = ? and user_id = ?", SqlVariable, UserId).Scan(&variable).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			fmt.Println("UpdateSqlVariableValidate query variables scan err = ", err)
			return false, err
		}
		return true, nil
	}

	return false, err

}
