package userbusiness

import (
	"context"
	"go-api/common"
	usermodel "go-api/module/user/model"
)

type FindUserStorage interface {
	Find(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type findUserBusiness struct {
	storage FindUserStorage
}

func NewFindUserBussiness(storage FindUserStorage) *findUserBusiness {
	return &findUserBusiness{storage: storage}
}

func (business *findUserBusiness) FindUser(context context.Context, conditions map[string]interface{}) (*usermodel.User, error) {
	user, err := business.storage.Find(context, conditions)

	if user == nil {
		if err == common.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
		}

		return nil, common.ErrDB(err)
	}

	user.Mask(false)

	return user, nil
}
