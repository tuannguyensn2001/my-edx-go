package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func NewStore(dbstring string) (*Store, error) {
	dsn := dbstring
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()

	return &Store{
		DB: db,
	}, err
}
