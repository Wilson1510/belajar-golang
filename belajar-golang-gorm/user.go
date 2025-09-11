package belajar_golang_gorm

import (
	"time"
	"gorm.io/gorm"
)

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
	Wallet Wallet `gorm:"foreignKey:user_id;references:id"`
	Addresses []Address `gorm:"foreignKey:user_id;references:id"`
	LikeProducts []Product `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:product_id"`
}

func (user *User) TableName() string {
	return "users"
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	if user.ID == "" {
		user.ID = "user-" + time.Now().Format("20060102150405")
	}
	return nil
}

type UserLogs struct {
	ID string `gorm:"primary_key;column:id;autoIncrement"`
	UserID string `gorm:"column:user_id"`
	Action string `gorm:"column:action"`
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"column:updated_at;autoCreateTime;autoUpdateTime:milli"`
}

func (userLogs *UserLogs) TableName() string {
	return "user_logs"
}