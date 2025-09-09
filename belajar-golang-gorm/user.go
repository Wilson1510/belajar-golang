package belajar_golang_gorm

import "time"

type Name struct {
	FirstName string
	MiddleName string
	LastName string
}

type User struct {
	ID string `gorm:"primary_key;column:id;<-:create"` // tidak perlu, ID adalah default primary key
	Name Name `gorm:"embedded"`
	Password string `gorm:"column:password"` // tidak perlu, semua kolom akan dibuat snake_case
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"` // autoCreateTime adalah default value untuk created_at
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"` // autoUpdateTime adalah default value untuk updated_at
	Information string `gorm:"-"`
}

func (user *User) TableName() string {
	return "users"
}