package userstorage

import (
	"context"
	usermodel "go-api/module/user/model"

	"github.com/google/uuid"
)

func (s *store) Delete(context context.Context, id uuid.UUID) error {
	tx := s.db.Begin()

	if err := tx.Table(usermodel.User{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
