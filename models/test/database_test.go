package models_test

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestConnect(t *testing.T) {
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	dsn := "root:rootroot@tcp(localhost:3306)/DEMO?charset=utf8&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB 연결에 실패하였습니다.")
	}

	//err = db.AutoMigrate(&Product{})
	//if err != nil {
	//	panic("DB 연결에 실패하였습니다.")
	//}
}
