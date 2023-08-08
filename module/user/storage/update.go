package userstorage

import (
	"context"
	"go-api/common"
	usermodel "go-api/module/user/model"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func (s *storage) Update(context context.Context, id uuid.UUID, data *usermodel.UserUpdate) (*usermodel.User, error) {
	var user usermodel.User
	if err := common.MappingModel(data, &user); err != nil {
		return nil, err
	}

	tx := s.db.Begin()

	if err := tx.Table(usermodel.TableName).Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}
