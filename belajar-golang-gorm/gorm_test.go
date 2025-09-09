package belajar_golang_gorm

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
)

func OpenConnection() *gorm.DB {
	password := os.Getenv("password")
	dialect := mysql.Open("root:"+password+"@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

var DB *gorm.DB = OpenConnection()

func TestGorm(t *testing.T) {
	fmt.Println(DB)
	assert.NotNil(t, DB)
}