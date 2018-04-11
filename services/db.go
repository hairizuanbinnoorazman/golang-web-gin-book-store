package services

import "github.com/jinzhu/gorm"

type Storage struct {
	DB *gorm.DB
}

func (s Storage) NewDB() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	s.DB = db
}
