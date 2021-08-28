package server_test

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestConnect(m *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

	db.AutoMigrate(&Product{})
}
