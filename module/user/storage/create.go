package userstorage

import (
	"context"
	usermodel "go-api/module/user/model"
)

func (s *storage) Create(context context.Context, data *usermodel.UserCreate) (*usermodel.User, error) {
	user := usermodel.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Phone:     data.Phone,
		Password:  data.Password,
	}

	tx := s.db.Begin()

	if err := tx.Table(usermodel.TableName).Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}
