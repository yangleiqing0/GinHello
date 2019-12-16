package model

import "fmt"

func SaveNameValidate(name, tableName string, userId int) (result bool, err error) {
	var count int
	err = db.Table(tableName).Where("name = ? and user_id = ?", name, userId).Count(&count).Error
	if err != nil {
		fmt.Println("count SaveNameValidate err = ", err)
		return
	}
	result = count == 0
	return
}

func UpdateNameValidate(name, tableName string, id, userId int) (result bool, err error) {
	var count int
	err = db.Table(tableName).Where("id != ? and name = ? and user_id = ?", id, name, userId).Count(&count).Error
	if err != nil {
		fmt.Println("count UpdateNameValidate err = ", err)
		return
	}
	return count == 0, nil
}
