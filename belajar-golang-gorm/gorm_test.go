package belajar_golang_gorm

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/stretchr/testify/assert"
)

func OpenConnection() *gorm.DB {
	password := os.Getenv("password")
	dialect := mysql.Open("root:"+password+"@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(3 * time.Hour)
	sqlDB.SetConnMaxIdleTime(3 * time.Minute)	
	return db
}

var db *gorm.DB = OpenConnection()

func TestGorm(t *testing.T) {
	fmt.Println(db)
	assert.NotNil(t, db)
}

func TestExecQuery(t *testing.T) {
	err := db.Exec("INSERT INTO sample (id, name) VALUES (?, ?)", 1, "test").Error
	assert.Nil(t, err)

	err = db.Exec("INSERT INTO sample (id, name) VALUES (?, ?)", 2, "test2").Error
	assert.Nil(t, err)

	err = db.Exec("INSERT INTO sample (id, name) VALUES (?, ?)", 3, "test3").Error
	assert.Nil(t, err)

	err = db.Exec("INSERT INTO sample (id, name) VALUES (?, ?)", 4, "test4").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id int
	Name string
}

func TestRawSql(t *testing.T) {
	var result Sample
	err := db.Raw("SELECT name FROM sample WHERE id = ?", 1).Scan(&result).Error
	assert.Nil(t, err)
	assert.Equal(t, "test", result.Name)

	var results []Sample
	err = db.Raw("SELECT name FROM sample").Scan(&results).Error
	fmt.Println(results)
	assert.Nil(t, err)
	assert.Equal(t, "test", results[0].Name)
	assert.Equal(t, 4, len(results))
}

func TestCreateUser(t *testing.T) {
	user := User{
		ID: "1",
		Name: Name{
			FirstName: "John",
			MiddleName: "Doe",
			LastName: "Doe",
		},
		Password: "123456",
		Information: "test",
	}
	result := db.Create(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, int64(1), result.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	users := []User{}
	for i := 0; i < 10; i++ {
		user := User{
			ID: strconv.Itoa(i+2),
			Name: Name{
				FirstName: "John"+strconv.Itoa(i), 
				MiddleName: "Doe"+strconv.Itoa(i), 
				LastName: "Wilson"+strconv.Itoa(i),
			},
			Password: "123456",
			Information: "test",
		}
		users = append(users, user)
	}
	result := db.Create(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, int64(10), result.RowsAffected)
}

func TestTransaction(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			ID: "12",
			Name: Name{FirstName: "John 12", MiddleName: "Doe 12", LastName: "Wilson 12"},
			Password: "123456",
			Information: "test",
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID: "13",
			Name: Name{FirstName: "John 13", MiddleName: "Doe 13", LastName: "Wilson 13"},
			Password: "123456",
			Information: "test",
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID: "14",
			Name: Name{FirstName: "John 14", MiddleName: "Doe 14", LastName: "Wilson 14"},
			Password: "123456",
			Information: "test",
		}).Error
		if err != nil {
			return err
		}
		return nil
	})
	assert.Nil(t, err)
}

func TestRollbackTransaction(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			ID: "15",
			Name: Name{FirstName: "John 15", MiddleName: "Doe 15", LastName: "Wilson 15"},
			Password: "123456",
			Information: "test",
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID: "12",
			Name: Name{FirstName: "John 12", MiddleName: "Doe 12", LastName: "Wilson 12"},
			Password: "123456",
			Information: "test",
		}).Error
		if err != nil {
			return err
		}
		return nil
	})
	assert.NotNil(t, err)
}

func TestCommitTransaction(t *testing.T) {
	tx := db.Begin()
	err := tx.Create(&User{
		ID: "16",
		Name: Name{FirstName: "John 16", MiddleName: "Doe 16", LastName: "Wilson 16"},
		Password: "123456",
		Information: "test",
	}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{
		ID: "17",
		Name: Name{FirstName: "John 17", MiddleName: "Doe 17", LastName: "Wilson 17"},
		Password: "123456",
		Information: "test",
	}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestQuery(t *testing.T) {
	var user User
	err := db.First(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID)

	user = User{}
	err = db.Last(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "9", user.ID)

	user = User{}
	err = db.Take(&user, "id = ?", "5").Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.ID)
}

func TestQueryAll(t *testing.T) {
	var users []User
	err := db.Find(&users, "id in (?)", []string{"1", "2", "3", "4"}).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func TestQueryCondition(t *testing.T) {
	var users []User
	err := db.Where("id in (?)", []string{"1", "2", "3", "4"}).Where("first_name = ?", "John").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))

	users = []User{}
	err = db.Where("id in (?)", []string{"1", "2", "3", "4"}).Or("first_name = ?", "John5").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))

	users = []User{}
	err = db.Not("first_name = ?", "John").Where("id in (?)", []string{"1", "2", "3", "4"}).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))
}

func TestSelect(t *testing.T) {
	var users []User
	err := db.Select("id", "first_name").Find(&users).Error
	assert.Nil(t, err)
	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
		assert.Equal(t, "", user.Name.MiddleName)
	}
	assert.Equal(t, 16, len(users))
}

func TestStructCondition(t *testing.T) {
	var users []User
	query := User{
		Name: Name{FirstName: "John", MiddleName: ""}, // middle name tidak berfungsi karena default value
	}
	err := db.Where(query).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestMapCondition(t *testing.T) {
	var users []User
	query := map[string]interface{}{
		"first_name": "",
	}
	err := db.Where(query).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, len(users))
}

func TestOrderLimitOffset(t *testing.T) {
	var users []User
	err := db.Order("id desc").Limit(3).Offset(4).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))
	assert.Equal(t, "5", users[0].ID)
	assert.Equal(t, "4", users[1].ID)
	assert.Equal(t, "3", users[2].ID)
}

type UserResponse struct {
	ID string
	FirstName string
	MiddleName string
	LastName string
}

func TestUseOtherStruct(t *testing.T) {
	var users []UserResponse
	err := db.Model(&User{}).Select("id", "first_name", "middle_name", "last_name").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 16, len(users))
	fmt.Println(users)
}

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