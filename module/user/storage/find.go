package userstorage

import (
	"context"
	usermodel "go-api/module/user/model"
)

func (s *store) FindDataWithConditions(context context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	var data usermodel.User

	if err := s.db.Table(usermodel.User{}.TableName()).Not("status", 0).Where(conditions).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
