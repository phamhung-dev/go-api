package userstorage

import (
	"context"
	"go-api/common"
	usermodel "go-api/module/user/model"
)

func (s *store) ListDataWithCondition(context context.Context, filter *usermodel.Filter, paging *common.Paging, moreKeys ...string) ([]usermodel.User, error) {
	var result []usermodel.User

	db := s.db.Table(usermodel.User{}.TableName()).Not("status", 0)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.Offset(offset).Limit(paging.Limit).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
