package models

import (
	"fmt"
	"go_gin_study/server"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	connectDatabase()
	//migrateTables()
}

func connectDatabase() {
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
		panic(err)
	}
}

func migrateTables() {
	if db.Migrator().HasTable(&User{}) == false {
		fmt.Println("migrate user table")
		err := db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&User{})
		if err != nil {
			panic(err)
		}
	}
}

// func CloseDB() {
// 	defer db.Close()
// }
