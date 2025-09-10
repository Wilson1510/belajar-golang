package belajar_golang_gorm

import (
	"fmt"
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
	"gorm.io/gorm"
)

func TestUpdate(t *testing.T) {
	var user User
	err := db.Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	user.Name.FirstName = "John 1"
	user.Name.MiddleName = "Doe 1"
	user.Name.LastName = "Wilson 1"
	err = db.Save(&user).Error
	assert.Nil(t, err)
}

func TestUpdate2(t *testing.T) {
	err := db.Where("id = ?", "1").Updates(User{
		Name: Name{
			FirstName: "Zacky",
			LastName: "Prasetio",
		},
	}).Error
	assert.Nil(t, err)

	err = db.Model(&User{}).Where("id = ?", "1").Updates(map[string]interface{}{
		"first_name": "Naufal",
		"last_name": "",
	}).Error
	assert.Nil(t, err)

	err = db.Model(&User{}).Where("id = ?", "1").Update("password", "7890").Error
	assert.Nil(t, err)
}

func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 10; i++ {
		userLogs := UserLogs{
			UserID: "1",
			Action: "login " + strconv.Itoa(i),
		}
		err := db.Create(&userLogs).Error
		assert.Nil(t, err)
		assert.NotEqual(t, 0, userLogs.ID)
		fmt.Println("User Logs ID: ", userLogs.ID)
	}
}

func TestCreateOrUpdate(t *testing.T) {
	userLogs := UserLogs{
		UserID: "1",
		Action: "login",
	}
	err := db.Save(&userLogs).Error
	assert.Nil(t, err)

	userLogs.Action = "logout"
	err = db.Save(&userLogs).Error
	assert.Nil(t, err)
}

func TestConflict(t *testing.T) {
	userLogs := User{
		ID: "41",
		Name: Name{
			FirstName: "Larry",
			MiddleName: "Parker",
			LastName: "Bird",
		},
		Password: "123456",
	}
	err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&userLogs).Error
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	err := db.Delete(&User{}, "id = ?", "41").Error
	assert.Nil(t, err)

	var user User
	err = db.Take(&user, "id = ?", "11").Error
	assert.Nil(t, err)
	err = db.Delete(&user).Error
	assert.Nil(t, err)

	err = db.Where("id = ?", "2").Delete(&User{}).Error
	assert.Nil(t, err)
}

func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserID: "1",
		Title: "Todo 1",
		Description: "Description 1",
	}
	err := db.Create(&todo).Error
	assert.Nil(t, err)

	err = db.Delete(&todo).Error
	assert.Nil(t, err)

	var todos []Todo
	err = db.Take(&todo, "id = ?", todo.ID).Error
	assert.Nil(t, err)
	assert.NotNil(t, todo.DeletedAt)
	assert.Equal(t, 0, len(todos))
}

func TestUnscoped(t *testing.T) {
	var todo Todo
	err := db.Unscoped().First(&todo, "id = ?", 1).Error
	assert.Nil(t, err)
	fmt.Println(todo)

	err = db.Unscoped().Delete(&todo).Error
	assert.Nil(t, err)

	var todos []Todo
	err = db.Unscoped().Find(&todos).Error
	assert.Nil(t, err)
	fmt.Println(todos)
}

func TestLock(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = ?", "1").Error
		assert.Nil(t, err)
		
		user.Name.FirstName = "kurniawan"
		err = tx.Save(&user).Error
		assert.Nil(t, err)
		return nil
	})
	assert.Nil(t, err)
}

func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		ID: "1",
		UserId: "1",
		Balance: 1000000,
	}
	err := db.Create(&wallet).Error
	assert.Nil(t, err)
}

func TestPreload(t *testing.T) {
	var user User
	err := db.Preload("Wallet").Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
	
	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)

	err = db.Joins("Wallet").Take(&user, "users.id = ?", "1").Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}

func TestAutoCreateUpdate(t *testing.T) {
	user := User{
		ID: "20",
		Name: Name{
			FirstName: "John",
			MiddleName: "Doe",
			LastName: "Wilson",
		},
		Wallet: Wallet{
			ID: "2",
			UserId: "2",
			Balance: 245000,
		},
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
}

func TestSkipCreateUpdate(t *testing.T) {
	user := User{
		ID: "21",
		Name: Name{
			FirstName: "John",
			MiddleName: "Doe",
			LastName: "Wilson",
		},
		Wallet: Wallet{
			ID: "3",
			UserId: "21",
			Balance: 345000,
		},
	}
	err := db.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}