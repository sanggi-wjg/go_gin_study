package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.nhnent.com/godo/cfo/server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var db *gorm.DB

func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Able color
		},
	)

	var err error
	db, err = gorm.Open(mysql.Open(getDSN()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	migrateTable()
	// db.SingularTable(true)
	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(100)
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseConfig.User,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Host,
		config.DatabaseConfig.Port,
		config.DatabaseConfig.DatabaseName,
	)
}

func migrateTable() {
	if db.Migrator().HasTable(&User{}) == false {
		err := db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&User{})
		if err != nil {
			panic(err)
		}
	}
}
