package services

import (
	"fmt"

	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Storage struct {
	DB *gorm.DB
}

func (s *Storage) NewDB() {
	if s.DB == nil {
		db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/hehe?charset=utf8&parseTime=True")

		if err != nil {
			fmt.Println(err.Error())
			panic("failed to connect database")
		}
		s.DB = db
	}
}

func (s *Storage) AutoMigrate() {
	s.DB.AutoMigrate(&models.User{})
}
