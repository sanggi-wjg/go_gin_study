package models

import (
	"fmt"
	"go_gin_study/server"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		server.DatabaseConfig.User,
		server.DatabaseConfig.Password,
		server.DatabaseConfig.Host,
		server.DatabaseConfig.Port,
		server.DatabaseConfig.DatabaseName,
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB 연결에 실패하였습니다.")
	}
}

// func CloseDB() {
// 	defer db.Close()
// }
