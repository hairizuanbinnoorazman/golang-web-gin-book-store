package users

import (
	"github.com/hairizuanbinnoorazman/golang-web-gin-book-store/models"
	"github.com/jinzhu/gorm"
)

type MySQLService struct {
	DB *gorm.DB
}

func (s MySQLService) GetByID(ID string) (models.User, error) {
	tempUser := models.User{ID: ID}
	result := s.DB.First(&tempUser)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return tempUser, nil
}

func (s MySQLService) Create(a *models.User) (models.User, error) {
	result := s.DB.Create(a)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return *a, nil
}

func (s MySQLService) Update(a *models.User) (models.User, error) {
	result := s.DB.Update(a)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return *a, nil
}

func (s MySQLService) GetByEmail(a string) (models.User, error) {
	tempUser := models.User{}
	result := s.DB.Where("email = ?", a).First(&tempUser)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return tempUser, nil
}
