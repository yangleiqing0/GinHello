package initDB

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DbInit() *gorm.DB{

	db := NewConn()

	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(10)

	// 自动迁移模式
	//db.AutoMigrate(&model.UserModel{})

	return db
}

func NewConn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/yang?charset=utf8&parseTime=true&loc=America%2FChicago")

	//defer func() {
	//	err := db.Close()
	//	if err != nil{
	//	    fmt.Println("db close err = ", err)
	//	}
	//}()

	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}
	return db
}