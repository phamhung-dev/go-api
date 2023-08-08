package userstorage

import (
	"context"
	"go-api/common"
	usermodel "go-api/module/user/model"
)

func (s *storage) List(context context.Context, filter *usermodel.Filter, paging *common.Paging, moreInfo ...string) ([]usermodel.User, error) {
	var users []usermodel.User

	db := s.db.Table(usermodel.TableName).Not("status", 0)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if paging.FakeCursor != "" {
		id, err := common.DecodeUID(paging.FakeCursor)

		if err != nil {
			return nil, err
		}

		db = db.Where("id < ?", id)
	} else {
		offset := (paging.Page - 1) * paging.Limit

		db = db.Offset(offset)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) > 0 {
		last := users[len(users)-1]

		paging.NextCursor = common.EncodeUID(last.ID)
	}

	return users, nil
}
