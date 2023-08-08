package userstorage

import (
	"context"
	"go-api/common"
	usermodel "go-api/module/user/model"

	"gorm.io/gorm"
)

func (s *storage) Find(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.TableName)

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Not("status", 0).Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}

	return &user, nil
}
