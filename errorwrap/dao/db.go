package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Dao db init
type Dao struct {
	db *gorm.DB
}

// User user information
type User struct {
	ID    int64
	Name  string
	Age   int8
	Sex   int8
	Phone string
	Err   error
}

// NewAutoDao return new db client
func NewAutoDao() (*Dao, error) {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})

	return &Dao{
		db: db,
	}, err

}

// QueryUser query user infomation
func (d *Dao) QueryUser() *User {
	var user User
	result := d.db.Find(&user)
	user.Err = result.Error
	return &user
}
