package belajar_golang_gorm

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
	"context"
	"gorm.io/gorm"
)

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

func TestCreateUserAddresses(t *testing.T) {
	user := User{
		ID: "3",
		Name: Name{
			FirstName: "fauzi",
			MiddleName: "bambang",
			LastName: "wijaya",
		},
		Wallet: Wallet{
			ID: "26",
			UserId: "26",
			Balance: 96000,
		},
		Addresses: []Address{
			{
				UserId: "3",
				Address: "Jl. Raya 1",
			},
			{
				UserId: "3",
				Address: "Jl. Raya 2",
			},
			{
				UserId: "3",
				Address: "Jl. Raya 3",
			},
		},
	}
	err := db.Save(&user).Error
	assert.Nil(t, err)
}

func TestGetUserAddresses(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").Take(&user, "users.id = ?", "25").Error
	assert.Nil(t, err)
	fmt.Println(user)

	var users []User
	err = db.Model(&User{}).Preload("Addresses").Joins("Wallet").Find(&users).Error
	assert.Nil(t, err)
	fmt.Println(users)
}

func TestBelongsToAddresses(t *testing.T) {
	var addresses []Address
	err := db.Preload("User").Find(&addresses).Error
	assert.Nil(t, err)
	fmt.Println(addresses)

	addresses = []Address{}
	err = db.Model(&Address{}).Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
	fmt.Println(addresses)
}

func TestBelongsToWallet(t *testing.T) {
	var wallets []Wallet
	err := db.Preload("User").Find(&wallets).Error
	assert.Nil(t, err)
	fmt.Println(wallets)

	wallets = []Wallet{}
	err = db.Model(&Wallet{}).Joins("User").Find(&wallets).Error
	assert.Nil(t, err)
	fmt.Println(wallets)
}

func TestManyToMany(t *testing.T) {
	product := Product{
		ID: "2",
		Name: "Product 2",
		Price: 250000,
	}
	err := db.Create(&product).Error
	assert.Nil(t, err)
	
	err = db.Table("user_like_product").Create([]map[string]interface{}{
		{
			"user_id": "21",
			"product_id": "2",
		},
	}).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create([]map[string]interface{}{
		{
			"user_id": "25",
			"product_id": "2",
		},
	}).Error
	assert.Nil(t, err)
}

func TestPreloadManyToMany(t *testing.T) {
	var product Product
	err := db.Preload("LikedByUsers").Take(&product, "id = ?", "2").Error
	assert.Nil(t, err)
	fmt.Println(product)

	var user User
	err = db.Preload("LikeProducts").Take(&user, "id = ?", "21").Error
	assert.Nil(t, err)
	fmt.Println(user)
}

func TestAssociationFind(t *testing.T) {
	var product Product
	err := db.Take(&product, "id = ?", "2").Error
	assert.Nil(t, err)
	assert.Equal(t, "Product 2", product.Name)

	var users []User
	err = db.Model(&product).Where("users.first_name = ?", "John").Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
	assert.Equal(t, "John", users[0].Name.FirstName)
}

func TestAssociationAppend(t *testing.T) {
	var product Product
	err := db.Take(&product, "id = ?", "2").Error
	assert.Nil(t, err)
	assert.Equal(t, "Product 2", product.Name)

	var user User
	err = db.Take(&user, "id = ?", "17").Error
	assert.Nil(t, err)
	assert.Equal(t, "John 17", user.Name.FirstName)

	err = db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)

}

func TestAssociationReplace(t *testing.T) {
	var wallet Wallet
	err := db.Take(&wallet, "id = ?", "2").Error
	assert.Nil(t, err)
	assert.Equal(t, int64(245000), wallet.Balance)

	var user User
	err = db.Take(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	err = db.Model(&user).Association("Wallet").Replace(&wallet)
	assert.NotNil(t, err)
	fmt.Println(err)
}

func TestAssociationDelete(t *testing.T) {
	var product Product
	err := db.Take(&product, "id = ?", "2").Error
	assert.Nil(t, err)
	assert.Equal(t, "Product 2", product.Name)

	var user User
	err = db.Take(&user, "id = ?", "17").Error
	assert.Nil(t, err)
	assert.Equal(t, "John 17", user.Name.FirstName)

	err = db.Model(&product).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)

}

func TestAssociationClear(t *testing.T) {
	var product Product
	err := db.Take(&product, "id = ?", "1").Error
	assert.Nil(t, err)
	assert.Equal(t, "Product 1", product.Name)

	err = db.Model(&product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
}

func TestPreloadCondition(t *testing.T) {
	var user User
	err := db.Preload("Wallet", "balance > ?", 100000).Take(&user, "id = ?", "3").Error
	assert.Nil(t, err)
	fmt.Println(user)
}

func TestNestedPreload(t *testing.T) {
	var wallet Wallet
	err := db.Preload("User.Addresses").Take(&wallet, "id = ?", "2").Error
	assert.Nil(t, err)

	fmt.Println(wallet)
	fmt.Println(wallet.User.Addresses)
}

func TestPreloadAll(t *testing.T) {
	var user User
	err := db.Preload(clause.Associations).Take(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	fmt.Println(user)
	fmt.Println(user.Addresses)
	fmt.Println(user.Wallet)
	fmt.Println(user.LikeProducts)
}

func TestJoinQuery(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 100000).Find(&users).Error
	assert.Nil(t, err)

	for _, user := range users {
		fmt.Println(user.Name.FirstName)
		fmt.Println(user.ID)
	}

	users = []User{}
	err = db.Joins("Wallet").Where("Wallet.balance > ?", 100000).Find(&users).Error
	assert.Nil(t, err)

	fmt.Println(len(users))
}

func TestCount(t *testing.T) {
	var count int64
	err := db.Model(&User{}).Joins("Wallet").Where("Wallet.balance > ?", 100000).Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(2), count)
}

type Result struct {
	SumBalance int64
	AvgBalance float64
	MaxBalance int64
	MinBalance int64
}

func TestAggregate(t *testing.T) {
	var result Result
	err := db.Model(&Wallet{}).Select(
		"sum(balance) as sum_balance, avg(balance) as avg_balance, max(balance) as max_balance, min(balance) as min_balance",
	).Take(&result).Error
	assert.Nil(t, err)
	fmt.Println(result.SumBalance)
	fmt.Println(result.AvgBalance)
	fmt.Println(result.MaxBalance)
	fmt.Println(result.MinBalance)
}

func TestGroupByHaving(t *testing.T) {
	var results []Result
	err := db.Model(&Wallet{}).Select(
		"sum(balance) as sum_balance, avg(balance) as avg_balance, max(balance) as max_balance, min(balance) as min_balance",
	).Joins("User").Group("user_id").Having("sum_balance > ?", 100000).Find(&results).Error
	assert.Nil(t, err)
	for _, result := range results {
		fmt.Println(result.SumBalance)
		fmt.Println(result.AvgBalance)
		fmt.Println(result.MaxBalance)
		fmt.Println(result.MinBalance)
		fmt.Println()
	}
}

func TestContext(t *testing.T) {
	var users []User
	ctx := context.Background()
	err := db.WithContext(ctx).Find(&users).Error
	assert.Nil(t, err)
	fmt.Println(len(users))
}

func BrokeWallet(db *gorm.DB) *gorm.DB {
	return db.Where("balance = ?", 0)
}

func SultanWallet(db *gorm.DB) *gorm.DB {
	return db.Where("balance > ?", 100000)
}

func TestScopes(t *testing.T) {
	var wallets []Wallet
	err := db.Scopes(BrokeWallet).Find(&wallets).Error
	assert.Nil(t, err)
	fmt.Println(len(wallets))

	wallets = []Wallet{}
	err = db.Scopes(SultanWallet).Find(&wallets).Error
	assert.Nil(t, err)
	fmt.Println(len(wallets))
}

func TestMigrating(t *testing.T) {
	err := db.Migrator().AutoMigrate(&GuestBook{})
	assert.Nil(t, err)
}

func TestHook(t *testing.T) {
	user := User{
		Name: Name{
			FirstName: "John",
			MiddleName: "Doe",
			LastName: "Wilson",
		},
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
	assert.NotEqual(t, "", user.ID)
	fmt.Println(user.ID)
}