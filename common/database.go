package common

import (
	"fmt"
	"ginclass/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	host := "localhost" //数据库地址
	port := "3306"      //数据库端口号
	dbname := "ginessential"
	username := "root"
	password := "123456"
	charset := "utf8mb4"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
		charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使用单数表名
		},
	})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	db.AutoMigrate(&model.User{})
	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}
