package userstorage

import (
	"context"
	usermodel "go-api/module/user/model"

	"github.com/google/uuid"
)

func (s *storage) Delete(context context.Context, id uuid.UUID) error {
	tx := s.db.Begin()

	if err := tx.Table(usermodel.TableName).Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
