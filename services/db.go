package services

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func (s Service) NewDB() Service {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	return Service{DB: db}
}
