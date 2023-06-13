package userstorage

import (
	"context"
	usermodel "go-api/module/user/model"
)

func (s *store) Create(context context.Context, data *usermodel.UserCreateRequest) error {
	tx := s.db.Begin()

	if err := tx.Table(usermodel.User{}.TableName()).Create(data).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
