package models_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestConnect(t *testing.T) {
	dsn := "root:rootroot@tcp(localhost:3306)/DEMO?charset=utf8&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	assert.Equal(t, nil, err)
}
