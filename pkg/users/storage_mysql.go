package users

import (
	"github.com/jinzhu/gorm"
)

type MySQLService struct {
	DB *gorm.DB
}

func (s MySQLService) GetByID(ID string) (User, error) {
	tempUser := User{ID: ID}
	result := s.DB.First(&tempUser)
	if result.Error != nil {
		return User{}, result.Error
	}
	return tempUser, nil
}

func (s MySQLService) Create(a *User) (User, error) {
	result := s.DB.Create(a)
	if result.Error != nil {
		return User{}, result.Error
	}
	return *a, nil
}

func (s MySQLService) Update(a *User) (User, error) {
	result := s.DB.Update(a)
	if result.Error != nil {
		return User{}, result.Error
	}
	return *a, nil
}

func (s MySQLService) GetByEmail(a string) (User, error) {
	tempUser := User{}
	result := s.DB.Where("email = ?", a).First(&tempUser)
	if result.Error != nil {
		return User{}, result.Error
	}
	return tempUser, nil
}
