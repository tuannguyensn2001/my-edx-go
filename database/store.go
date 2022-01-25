package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func NewStore() (*Store, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/edx?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()

	return &Store{
		DB: db,
	}, err
}
